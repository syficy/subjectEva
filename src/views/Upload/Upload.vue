<template>
    <div>
    <nav-bar>
        <div slot="center">
            提交评价
        </div>
    </nav-bar>
        <div>
            <el-input placeholder="请输入内容" v-model="inputClass">
                <template slot="prepend">课程名称</template>
            </el-input>
        </div>
        <div>
            <el-input placeholder="请输入内容" v-model="inputTea">
                <template slot="prepend">教师名称</template>
            </el-input>
        </div>
        <el-input
                type="textarea"
                :rows="7"
                placeholder="请输入评价内容"
                v-model="inputText">
        </el-input>

        <div class="block" style="margin-top: 20%;">
            <span class="demonstration">评分</span>
            <el-rate v-model="rate" ></el-rate>
        </div>


        <div style="margin-bottom:80px; position: fixed;bottom: 0;right:10%">
        <el-button type="success" icon="el-icon-check" size="medium" circle @click="dialogTableVisible = true"></el-button>
        </div>

        <el-dialog title="确认提交" :visible.sync="dialogTableVisible" width="80%" style="margin-top: 20%">
            <el-form label-position="right" label-width="80px"  >
                <el-input placeholder="请输入学号" v-model="inputstu">
                    <template slot="prepend">学号</template>
                </el-input>
                <el-input placeholder="请输入密码" v-model="inputpass">
                    <template slot="prepend">密码</template>
                </el-input>
            </el-form>
            <el-button type="primary" style="width:40%" @click="upload()">确认提交</el-button>
        </el-dialog>


    </div>
</template>

<script>
    import NavBar from "../../components/common/navbar/NavBar";
    import {insertRequest,searchTeacherRequest,querySubjectRequest} from "../../network/eva";

    import VueMarkdown from 'vue-markdown'

    export default {
        name: "Upload.vue",
        components:{
            NavBar,
            VueMarkdown
        },
        data() {
            return {
                inputText: '',
                inputClass:"",
                inputTea:"",
                inputstu: "",
                inputpass: "",
                rate:3,
                dialogTableVisible:false,
                formLabelAlign: {
                    schoolid: 0,
                    password: '',
                },

                options: [],
                value: '',
                count: 10,

                subjectData:[]
            }
        },
        mounted() {
            querySubjectRequest().then(res=>{
                window.console.log(res.code)
                if (res.code==200){
                    window.console.log("查询subject成功")
                }
                this.subjectData=res.data
                var len = this.subjectData.length
                for(var i=0;i<len;i++)
                {
                    let data = {
                        "value":this.subjectData[i],
                        "label":this.subjectData[i]
                    }
                    this.options.push(data)
                }
                window.console.log(this.options)
            })
        },
        methods:{
            upload(){
                if(this.inputClass.length <= 1 ||this.inputTea.length < 3)
                {
                    return this.$message.error("课程或教师字符长度错误")
                }
                if(this.inputText.length < 10 ||this.inputText.length > 140)
                {
                    return this.$message.error("评价字符应为10-140")
                }


                let param = {
                    "subject": this.inputClass,
                    "teacher": this.inputTea,
                    "studentid": parseInt(this.inputstu),
                    "rate": this.rate,
                    "evtext": this.inputText
                }
                window.console.log(param)
                insertRequest(param).then(res=>{
                    if (res.code==200){
                        this.dialogTableVisible=false
                        return this.$message.success("上传成功")
                    }
                })
            }
        }
    }
</script>

<style scoped>
    .el-input{
        margin-top: 5%;
        margin-bottom: 5%;
    }
    .el-button{
        width: 50px;
        height:50px;
    }
</style>