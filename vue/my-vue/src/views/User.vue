<template>
    <div class="user-header">
        <el-button type="primary" @click="dialogVisible = true">+新增</el-button>
        <el-form :inline="true" :model="formInline" >
            <el-form-item label="请输入">
                <el-input v-model="formInline.keyword" placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="handleSearch">搜索</el-button>
            </el-form-item>
        </el-form>

        <el-dialog
            v-model="dialogVisible"
            title="新增用户"
            width="35%"
            :before-close="handleClose"
        >
        <el-form :inline="true" :model="formUser" >
            <el-row>
                <el-col :span="12">
                    <el-form-item label="姓名">
                        <el-input v-model="formUser.name" placeholder="请输入姓名" />
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="年龄">
                        <el-input v-model="formUser.age" placeholder="请输入年龄" />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="性别">
                        <el-select
                        v-model="formUser.sex"
                        placeholder="请选择性别"
                        clearable
                        >
                        <el-option label="男" value="1" />
                        <el-option label="女" value="0" />
                    </el-select>
                </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="出生日期">
                        <el-date-picker
                            v-model="formUser.birth"
                            type="date"
                            placeholder="Pick a date"
                            style="width: 100%"
                            />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="地址">
                        <el-input v-model="formUser.addr" placeholder="请输入地址" />
                    </el-form-item>
                </el-col>
            </el-row>
            
        </el-form>
            <template #footer>
            <div class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="onSubmit">
                    确认
                </el-button>
            </div>
            </template>
        </el-dialog>

    </div>
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
        <el-pagination
            small
            background
            layout="prev, pager, next"
            :total="config.total"
            @current-change = "changePage"
            class="pager mt-4"
        />
    </div>
</template>

<script>
import { defineComponent, getCurrentInstance, onMounted, reactive, ref } from 'vue';
import { ElMessageBox } from 'element-plus'
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

        const config = reactive({
            total: 10,
            page: 1,
            name: "",
        })

        const formInline = reactive({
            keyword: "",
        })

        onMounted(()=>{
            getUserData(config);
        })

        const changePage = (page)=>{
            config.page = page;
            getUserData(config)
        };

        const getUserData = async (config)=>{
            let res = await proxy.$api.getUserData();
            // let res = await proxy.$api.getUserData(config); // 会报网络错误
            config.total = res.count;
            // list.value =res.list
            list.value = res.list.map((item)=>{
                item.sexLabel = item.sex === 1 ? '男' : '女';
                return item;
            });
        };

        const handleSearch = ()=>{
            config.name = formInline.keyword;
            getUserData(config);
        };

        const dialogVisible = ref(false);

        const handleClose = (done) => {
        ElMessageBox.confirm('确定关闭吗?')
            .then(() => {
                done()
            })
            .catch(() => {
                // catch error
            })
        }

        const formUser = reactive({
            name:'',
            age:'',
            sex:'',
            birth:'',
            addr:'',
        })

        const onSubmit = async ()=>{
            let res = await proxy.$api.addUser(formUser)
            console.log(res)

            getUserData(config)
        }

        return {
            list,
            tableLabel,
            config,
            changePage,
            formInline,
            handleSearch,
            handleClose,
            dialogVisible,
            formUser,
            onSubmit,
        }
    }
})
</script>


<style lang="less" scoped>

.user-header{
    display: flex;
    justify-content: space-between;
}

.table{
    position: relative;
    height: 520px;
    .pager{
        position: absolute;
        right: 0;
        bottom: -20px;
    }
}



</style>