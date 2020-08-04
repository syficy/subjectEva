<template>
    <div>
        <nav-bar>
            <div slot="center">
                我的
            </div>
        </nav-bar>


        <div >
            <!--头像部分-->
            <div class="avatar-wrapper">
                <div class="avatar">
                    <img src="../../assets/images/logo.png" style="width: 7.5rem;height:7.5rem">
                </div>
            </div>


            <!--注册-->

            <el-radio-group v-model="labelFunc" size="small">
                <el-radio-button label="log">登入</el-radio-button>
                <el-radio-button label="register">注册</el-radio-button>
            </el-radio-group>
            <div style="margin: 20px;"></div>
            <el-form label-position="right" label-width="80px" style="margin-right:10%"  :model="formLabelAlign" v-show="isloged==false">
                <el-form-item label="学号" >
                    <el-input placeholder="请输入学号" v-model="formLabelAlign.schoolid"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input placeholder="请输入密码" v-model="formLabelAlign.password" show-password></el-input>
                </el-form-item>
                <el-form-item label="确认密码" v-show="labelFunc=='register'">
                    <el-input placeholder="请确认密码" v-model="formLabelAlign.passwordConf" show-password></el-input>
                </el-form-item>
                <el-form-item label="昵称" v-show="labelFunc=='register'">
                    <el-input placeholder="请输入昵称" v-model="formLabelAlign.name"></el-input>
                </el-form-item>
                <el-form-item label="电话号码" v-show="labelFunc=='register'">
                    <el-input placeholder="请输入电话号码" v-model="formLabelAlign.phonenumber"></el-input>
                </el-form-item>
            </el-form>
            <el-button v-show="labelFunc=='log' && isloged==false"  @click="studentLogin()">登入</el-button>
            <el-button v-show="labelFunc=='register' && isloged==false"  @click="studentRegister()">注册</el-button>


            <el-collapse v-model="activeName" accordion v-show="isloged" style="margin-left: 10%;margin-right: 10%">
                <el-collapse-item title="修改密码" name='chpd'>
                    <div>
                        <el-card shadow="never" v-show="activeName=='chpd' ">
                            <div  style="margin-top: 2%;">
                                <el-input style="margin-top: 5%;" placeholder="请输入原密码" v-model="chgPdOri" show-password>
                                    <template slot="prepend">原密码</template>
                                </el-input>
                                <el-input style="margin-top: 5%;" placeholder="请输入新密码" v-model="chgPdNew" show-password>
                                    <template slot="prepend">新密码</template>
                                </el-input>
                                <el-input style="margin-top: 5%;"  placeholder="请确认新密码" v-model="chgPdNewConf" show-password>
                                    <template slot="prepend">新密码</template>
                                </el-input>
                                <el-button style="margin-top: 10%;" type="success" round @click="chgPd()">确认修改密码</el-button>
                            </div>
                        </el-card>
                    </div>
                </el-collapse-item>
                <el-collapse-item title="修改信息" name='chinfo'>
                    <div>
                        <el-card shadow="never"  v-show="activeName=='chinfo'">
                            <div style="margin-top: 2%;">
                                <el-input style="margin-top: 5%;" placeholder="请输入新昵称" v-model="chgInfoName">
                                    <template slot="prepend">新昵称</template>
                                </el-input>
                                <el-input style="margin-top: 5%;" placeholder="请输入新号码" v-model="chgInfoPhone" >
                                    <template slot="prepend">新号码</template>
                                </el-input>
                                <el-button style="margin-top: 10%;" type="success" round @click="chgInfo()">确认修改信息</el-button>
                            </div>
                        </el-card>
                    </div>
                </el-collapse-item>
            </el-collapse>







<!--
            &lt;!&ndash;退出&ndash;&gt;
            <el-popconfirm
                    confirmButtonText='好的'
                    cancelButtonText='不用了'
                    icon="el-icon-info"
                    iconColor="red"
                    title="这是一段内容确定删除吗？"
            >
                <el-button slot="reference">登出</el-button>
            </el-popconfirm>-->
        </div>




    </div>
</template>

<script>
    import NavBar from "../../components/common/navbar/NavBar";
    import {loginRequest,registerRequest,changePdRequest,changeInfoRequest} from "../../network/user.js"
    export default {
        name: "Profile.vue",
        components:{
            NavBar
        },
        data() {
            return {
                pw: "",
                zh: "",
                isloged: false,
                labelFunc: "log",
                formLabelAlign: {
                    schoolid: '',
                    password: '',
                    name: '',
                    passwordConf: '',
                    phonenumber: ''
                },


                activeName: 0,

                /*修改密码*/
                chgPdOri: '',
                chgPdNew: '',
                chgPdNewConf: '',


                /*修改信息*/
                chgInfoName: "",
                chgInfoPhone: "",


            }
        },

        methods:{
            studentLogin(){
                if(this.formLabelAlign.schoolid.length!=13)
                {
                    return this.$message.error("学号应为13位数字")
                }


                window.console.log("hello")
                let param = {
                    "schoolid": parseInt(this.formLabelAlign.schoolid),
                    "password": this.formLabelAlign.password
                }

                loginRequest(param).then(res=>{
                    if (res.code==200){
                        this.isloged=true
                        return this.$message.success("登入成功")
                    }
                    else if (res.code==401){
                        return this.$message.error("用户未注册，请先注册")
                    }
                    else if (res.code==402){
                        return this.$message.error("密码错误")
                    }
                }).catch(
                )
            },

            studentRegister(){
                if(this.formLabelAlign.schoolid.length!=13)
                {
                    return this.$message.error("学号应为13位数字")
                }
                if(this.formLabelAlign.password.length > 16 ||this.formLabelAlign.password.length < 6)
                {
                    return this.$message.error("密码应为6-16位字符")
                }
                if(this.formLabelAlign.phonenumber.length!=11)
                {
                    return this.$message.error("电话号码应为11位数字")
                }
                if(this.formLabelAlign.name=="" || this.formLabelAlign.password=="")
                {
                    return this.$message.error("必填项不能为空")
                }




                if(this.formLabelAlign.password!=this.formLabelAlign.passwordConf){
                    return this.$message.error("两次密码不相同")
                }

                let param = {
                    "schoolid": parseInt(this.formLabelAlign.schoolid),
                    "password": this.formLabelAlign.password,
                    "name": this.formLabelAlign.name,
                    "phonenumber":  parseInt(this.formLabelAlign.phonenumber)
                }
                registerRequest(param).then(res=>{
                    if (res.code==200){
                        this.labelFunc="log"
                        return this.$message.success("注册成功")
                    }
                    else if(res.code==401) {
                        return this.$message.error("用户已存在，请登入，如忘记密码请联系qq512858048")
                    }
                    else{
                        return this.$message.error("未知错误，请联系qq512858048")
                    }
                })
            },

            chgPd(){

                if(this.chgPdNew.length > 16 ||this.chgPdNew.length < 6)
                {
                    return this.$message.error("密码应为6-16位字符")
                }


                if(this.chgPdNewConf!=this.chgPdNew){
                    return this.$message.error("两次密码不相同")
                }
                window.console.log(this.formLabelAlign)
                let param = {
                    "schoolid": parseInt(this.formLabelAlign.schoolid),
                    "oldpassword": this.chgPdOri,
                    "password": this.chgPdNew,
                }

                changePdRequest(param).then(res=>{
                    if (res.code==200){
                        this.labelFunc="log"
                        return this.$message.success("修改成功")
                    }
                    else if(res.code==401) {
                        return this.$message.error("原来密码错误，如忘记密码请联系qq512858048")
                    }
                    else{
                        return this.$message.error("未知错误，请联系qq512858048")
                    }
                })


            },

            chgInfo(){
                if(this.chgInfoPhone.length!=11)
                {
                    return this.$message.error("电话号码应为11位数字")
                }
                if(this.chgInfoName=="")
                {
                    return this.$message.error("必填项不能为空")
                }



                let param = {
                    "schoolid": parseInt(this.formLabelAlign.schoolid),
                    "name": this.chgInfoName,
                    "phonenumber": parseInt(this.chgInfoPhone)
                }

                changeInfoRequest(param).then(res=>{
                    if (res.code==200){
                        this.labelFunc="log"
                        return this.$message.success("修改成功")
                    }
                    else if(res.code==401) {
                        return this.$message.error("密码错误，如忘记密码请联系qq512858048")
                    }
                    else{
                        return this.$message.error("未知错误，请联系qq512858048")
                    }
                })


            }

        }

    }
</script>

<style scoped>
    .avatar-wrapper {
        -webkit-box-orient: vertical;
        -webkit-box-direction: normal;
        -ms-flex-direction: column;
        flex-direction: column;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;
        padding: 2.5rem;
    }

    .avatar {
        -webkit-box-flex: 0;
        -ms-flex: 0 0 auto;
        flex: 0 0 auto;
        width: 7.5rem;
        height: 7.5rem;
        background-color: #f9f9f9;
        border-radius: 50%
    }

    .avatar:hover {
        -webkit-transform: rotate(666turn);
        transform: rotate(666turn);
        transition-delay: 1s;
        transition-property: all;
        transition-duration: 59s;
        transition-timing-function: cubic-bezier(.34, 0, .84, 1)
    }

    .avatar {
        display: inline-block;
        position: relative;
        background-position: 50%;
        background-size: cover;
        background-repeat: no-repeat;
        background-color: #eee;
    }
</style>