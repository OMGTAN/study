<template>
    <div class="table">
        <el-table :data="list" style="width: 100% " height= 500>
            <el-table-column v-for=" item in tableLabel" :key="item.prop" :label="item.label" :prop="item.prop" :width="item.width ? item.width : 125"/>
            <el-table-column fixed="right" label="操作" min-width="180">
            <template #default>
                <el-button size="small" 
                >编辑</el-button
                >
                <el-button  type="danger" size="small">删除</el-button>
            </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import { defineComponent, getCurrentInstance, onMounted, reactive, ref } from 'vue';
export default  defineComponent({
    setup(){
        const {proxy} = getCurrentInstance();

        const tableLabel = reactive([
            {
                prop: "name",
                label: "姓名",
            },
            {
                prop: "age",
                label: "年龄",
            },
            {
                prop: "sexLabel",
                label: "性别",
            },
            {
                prop: "birth",
                label: "出生日期",
                width: 200,
            },
            {
                prop: "addr",
                label: "地址",
                width: 320,
            },
        ])

        let list = ref([])
        onMounted(()=>{
            getUserData();
        })

        const getUserData = async ()=>{
            let res = await proxy.$api.getUserData();
            console.log(res)
            // list.value =res.list
            list.value =res.list.map((item)=>{
                item.sexLabel = item.sex === 1 ? '男' : '女'
                return item
            })
        }

        return {
            list,
            tableLabel,
        }
    }
})
</script>