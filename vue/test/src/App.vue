<template>
    <div id="app">
        count is {{ count }}
        <br>
        <input type="button" value="test" size="11 " @click="handleBatchDownload">
    </div>
</template>

<script>
    //jszip是一个用于创建、读取和编辑.zip文件的JavaScript库 可以将文件或者图片打包成一个Zip文件。
    import JSZip from 'jszip'
    import FileSaver from 'file-saver'
    import axios from 'axios'

    export default ({
        data() {
            return { count: 0 }
        },

        methods: {
    // 重写请求 下载文件
    getFile(url){
        // let that = this;
        return new Promise((resolve, reject) => {
            axios({
                method:'get',
                url,
                responseType: 'blob'
            }).then(data => {
                resolve(data.data)
            }).catch(error => {
                // 在此写入你的错误处理
                // 可以在此处处理promise.all失败的情况，放在methods里主要也是为了方便对data里的变量处理
                reject(error.toString())
            })
        })
    },
    
    // 打压缩包下载
    handleBatchDownload() {
        const _this = this;
        const dataUrl = [];
        // for (let file of _this.fileTable) {
            let url = 'http://localhost:8080/file/downLoad'
            dataUrl.push(url)	// 把所有要打包文件的url放进数组
        // }
        const zip = new JSZip()	// 创建一个JSZip实例
        const cache = {}
        const promises = []
        dataUrl.forEach(item => {
            const promise = _this.getFile(item).then(data => { // 下载文件, 并存成ArrayBuffer对象
                const arr_name = item.split("/")  // stringObject.split 将字符串分割成字符串数组
                const file_name = arr_name[arr_name.length - 1] // 获取文件名（数组最后一项）
                zip.file("aaa.txt", data, { binary: true }) // 逐个添加文件
                cache[file_name] = data
            })
            promises.push(promise)
        })
        Promise.all(promises).then(() => {
            zip.generateAsync({type:"blob"}).then(content => { // 生成二进制流
                // 利用file-saver保存文件 _this.fileTableTitle是我压缩包名字
                FileSaver.saveAs(content, "tfsadf" + ".zip") 
            })
        })

    }
}
    })



</script>