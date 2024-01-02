<template>
  <div>
    <!-- <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
      <el-radio-button :label="false">expand</el-radio-button>
      <el-radio-button :label="true">collapse</el-radio-button>
    </el-radio-group> -->
  <el-menu
    default-active="2"
    class="el-menu-vertical-demo"
    :collapse="isCollapse"
    @open="handleOpen"
    @close="handleClose"
  >
    <el-menu-item index="2">
      <el-icon><icon-menu /></el-icon>
      <template #title>Navigator Two</template>
    </el-menu-item>

      <el-sub-menu index="1-4">
        <el-icon><icon-menu /></el-icon>
        <template #title>Navigator One</template>
        
        <el-menu-item index="1-4-1">item one</el-menu-item>
      </el-sub-menu>

  </el-menu>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'
import { fa } from 'element-plus/es/locale';

const isCollapse = ref(false)
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const menuData = [
  {
    path: "/",
    name: "home",
    label: "首页",
    icon: "s-home",
    url: "Home/Home",
  },
  {
    path: "/mall",
    name: "mall",
    label: "商品管理",
    icon: "video-play",
    url: "MallManage/MallManage",
  },
  {
    path: "/user",
    name: "user",
    label: "用户管理",
    icon: "user",
    url: "UserManage/UserManage",
  },
  {
    label: "其他",
    icon: "location",
    children: [
      {
        path: "/page1",
        name: "page1",
        label: "页面1",
        icon: "setting",
        url: "Other/PageOne",
      },
      {
        path: "/page2",
        name: "page2",
        label: "页面2",
        icon: "setting",
        url: "Other/PageTwo",
      }
    ]
  }
]

</script>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
