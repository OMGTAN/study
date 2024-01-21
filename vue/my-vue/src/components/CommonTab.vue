<template>
    <div class="tags">
        <el-tag :key="tag.name" v-for="(tag, index) in tabList" :closable="tag.name !== 'home'" :disable-transitions="false" 
            :effect="(route.name === tag.name) ? 'dark' : 'plain'" @click="clickMenu(tag)">
            {{ tag.label }}
        </el-tag>
    </div>
</template>

<script>
    import {useRouter}  from 'vue-router';
    import { useStore } from 'vuex';

    export default{
        setup(){

            const route = useRouter()
            const store = useStore()

            const tabList = store.state.tabList

            const clickMenu = (item)=>{
                console.log(route.currentRoute.value)
                // console.log(item.name)
                route.push(item.name)

            }

            return {
                tabList,
                clickMenu,
                route,
            }

        }
    }

</script>

<style lang="less" scoped>

    .tags{
        width: 100%;
        padding: 20px 20px 0px 20px;
        .el-tag{
            margin-right: 20px;
            cursor: pointer;
        }
    }

</style>