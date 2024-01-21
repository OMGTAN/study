<template>
    <div class="tags">
        <el-tag :key="tag.name" v-for="(tag, index) in tabList" :closable="tag.name !== 'home'" :disable-transitions="false" 
            :effect="($route.name === tag.name) ? 'dark' : 'plain'" @click="clickMenu(tag)" @close="closeTab(tag, index)">
            {{ tag.label }}
        </el-tag>
    </div>
</template>

<script>
    import {useRouter, useRoute}  from 'vue-router';
    import { useStore } from 'vuex';

    export default{
        setup(){

            const router = useRouter()
            const route = useRoute()
            const store = useStore()

            const tabList = store.state.tabList

            const clickMenu = (item)=>{
                // console.log(item.name)
                router.push({name: item.name})

            }

            const closeTab = (tag, index)=>{
                let length = tabList.length -1;
                store.commit("closeTab", tag)

                if(route.name !== tag.name){
                    return
                }

                if(index === length){
                    router.push({name: tabList[index-1].name})
                }else{
                    router.push({name: tabList[index].name})
                }
            }

            return {
                tabList,
                clickMenu,
                closeTab,
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