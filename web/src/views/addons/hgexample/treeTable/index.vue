<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="表格例子"> 一个树形表格例子 </n-card>
    </div>
    <n-grid cols="1 s:1 m:1 l:4 xl:4 2xl:4" responsive="screen" :x-gap="12">
      <n-gi span="1">
        <n-card :segmented="{ content: true }" :bordered="false" size="small" class="proCard">
          <template #header>
            <n-space>
              <n-button type="info" icon-placement="left" @click="openCreateDrawer">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <PlusOutlined />
                    </n-icon>
                  </div>
                </template>
                添加
              </n-button>
              <n-button
                type="info"
                icon-placement="left"
                @click="openEditDrawer"
                :disabled="!formParams?.id || formParams?.id < 1"
              >
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <EditOutlined />
                    </n-icon>
                  </div>
                </template>
                编辑
              </n-button>
              <n-button type="error" icon-placement="left" @click="handleDel">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <DeleteOutlined />
                    </n-icon>
                  </div>
                </template>
                删除
              </n-button>
              <n-button type="info" icon-placement="left" @click="packHandle">
                {{ expandedKeys.length ? '收起' : '展开' }}
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <AlignLeftOutlined />
                    </n-icon>
                  </div>
                </template>
              </n-button>
            </n-space>
          </template>
          <div class="w-full menu">
            <n-input v-model:value="pattern" placeholder="输入地区名称搜索">
              <template #suffix>
                <n-icon size="18" class="cursor-pointer">
                  <SearchOutlined />
                </n-icon>
              </template>
            </n-input>
            <div class="py-3 menu-list">
              <template v-if="loading">
                <div class="flex items-center justify-center py-4">
                  <n-spin size="medium" />
                </div>
              </template>
              <template v-else>
                <n-tree
                  block-line
                  cascade
                  checkable
                  :virtual-scroll="true"
                  :pattern="pattern"
                  :data="treeData"
                  :expandedKeys="expandedKeys"
                  style="height: 75vh"
                  @update:selected-keys="selectedTree"
                  @update:expanded-keys="onExpandedKeys"
                />
              </template>
            </div>
          </div>
        </n-card>
      </n-gi>
      <n-gi span="3">
        <n-card :segmented="{ content: true }" :bordered="false" size="small" class="proCard">
          <template #header>
            <n-space>
              <n-icon size="18">
                <FormOutlined />
              </n-icon>
              <span>正在编辑 {{ treeItemTitle ? `：${treeItemTitle}` : '' }}</span>
              <span style="font-size: 14px">{{
                treeItemTitle ? '' : '从列表选择一项后，进行编辑'
              }}</span>
            </n-space>
          </template>

          <List
            :checkedId="checkedId"
            :optionTreeData="optionTreeData"
            @reload-table="reloadTable"
          />
        </n-card>
      </n-gi>
    </n-grid>

    <Edit
      @reload-table="reloadTable"
      @update-show-modal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
      :optionTreeData="optionTreeData"
    />
  </div>
</template>
<script lang="ts" setup>
  import { onMounted, ref, unref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import {
    AlignLeftOutlined,
    FormOutlined,
    PlusOutlined,
    EditOutlined,
    SearchOutlined,
    DeleteOutlined,
  } from '@vicons/antd';
  import List from './list.vue';
  import { Delete, Select } from '@/api/addons/hgexample/treeTable';
  import Edit from './edit.vue';
  import { newState, loadOptions } from './model';

  const showModal = ref(false);
  const message = useMessage();
  const dialog = useDialog();
  let treeItemKey = ref([]);
  let expandedKeys = ref([]);
  const treeData = ref([]);
  const loading = ref(true);
  const treeItemTitle = ref('');
  const checkedId = ref(0);
  const pattern = ref('');
  const optionTreeData = ref<any>([]);
  const formParams = ref(newState(null));

  function openCreateDrawer() {
    showModal.value = true;
    formParams.value = newState(null);
  }

  function openEditDrawer() {
    showModal.value = true;
  }

  function selectedTree(keys, opts) {
    if (keys.length) {
      const treeItem = opts[0];
      treeItemKey.value = keys;
      treeItemTitle.value = treeItem.label;
      formParams.value = newState(treeItem);
      checkedId.value = treeItem.id;
    } else {
      treeItemKey.value = [];
      treeItemTitle.value = '';
    }
  }

  function handleDel() {
    dialog.info({
      title: '提示',
      content: `您确定想删除吗?`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ ...formParams.value }).then(async (_res) => {
          message.success('操作成功');
          await reloadTable();
        });
      },
      onNegativeClick: () => {
        message.error('已取消');
      },
    });
  }

  function packHandle() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = unref(treeData).map((item: any) => item.key as string) as [];
    }
  }

  onMounted(async () => {
    await reloadTable();
  });

  async function reloadTable() {
    const treeMenuList = await Select();
    Object.assign(
      formParams,
      treeMenuList.list.map((item) => item.key)
    );
    treeData.value = [];
    optionTreeData.value = [
      {
        id: 0,
        key: 0,
        label: '根目录',
        pid: 0,
        title: '根目录',
      },
    ];
    treeData.value = treeMenuList.list;
    optionTreeData.value = optionTreeData.value.concat(treeMenuList.list);
    loading.value = false;
  }

  function onExpandedKeys(keys) {
    expandedKeys.value = keys;
  }

  function updateShowModal(value) {
    showModal.value = value;
  }

  onMounted(() => {
    loadOptions();
  });
</script>
