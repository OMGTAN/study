import {createStore} from 'vuex'

export default createStore({
    state:{
        isCollapse: true,
        currentMenu: null,
        tabList:[
            {
                path: "/",
                name: "home",
                label: "首页",
                icon: "s-home",
            }
        ]
    },
    mutations:{
        updateIsCollapse(state, payload){
            state.isCollapse=!state.isCollapse
        },
        selectMenu(state, val){
            // val.name == 'home' ? (state.currentMenu = null) : (state.currentMenu = val);
            if(val.name == 'home'){
                state.currentMenu = null
            }else{
                state.currentMenu = val
                let res = state.tabList.findIndex((item)=>{
                    return item.name == val.name
                })
                res == -1 ? state.tabList.push(val): ''
            }
        },
    }
})