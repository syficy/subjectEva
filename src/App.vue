<template>
  <div id="app">
    <!--路由占位符-->
    <div>
    <router-view></router-view>
    </div>
    <div style="float:bottom">
    <MainTableBar v-show="istabBarshow"></MainTableBar>
    </div>
  </div>
</template>



<script>
import MainTableBar from "./components/content/MainTableBar";
export default {
  name:'App',
  components:{
    MainTableBar
  },
  data(){
    return{
      istabBarshow:true
    }
  }
}

</script>



<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
