<template>
    <tab-bar>
        <tab-bar-item path="/home" >
            <i slot="item-icon" class="el-icon-house"></i>
            <i slot="item-icon-active" class="el-icon-s-home"></i>
            <div slot="item-text">主页</div>
        </tab-bar-item>
        <tab-bar-item path="/Search" >
            <i slot="item-icon" class="el-icon-search"></i>
            <i slot="item-icon-active" class="el-icon-star-on"></i>
            <div slot="item-text">搜索</div>
        </tab-bar-item>
        <tab-bar-item path="/Upload">
            <i slot="item-icon" class="el-icon-upload2"></i>
            <i slot="item-icon-active" class="el-icon-upload"></i>
            <div slot="item-text">提交</div>
        </tab-bar-item>
        <tab-bar-item path="/profile">
            <i slot="item-icon" class="el-icon-user"></i>
            <i slot="item-icon-active" class="el-icon-user-solid"></i>
            <div slot="item-text">账号</div>
        </tab-bar-item>
    </tab-bar>
</template>

<script>
    import TabBarItem from "../common/tabbar/TabBarItem";
    import TabBar from "../common/tabbar/TabBar";
    export default {
        name: "MainTableBar",
        components: {TabBarItem,TabBar},
    }
</script>

<style scoped>

</style>