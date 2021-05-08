<template>
  <div id="app">
    <div class="logo" >
      <img style="width: 170px;" src="@/assets/images/logo_2.png" class="image">
    </div>
    <router-view/>
  </div>
</template>
<script type="es6">
import router from '@/router'
export default {
  methods: {
    link: function () {
      router.push({
        name: 'Home'
      })
    }
  }
}
</script>
<style>
html, body, #app {
  background-color: #2a2c3b;
  font-size: 14px;
  height: 100%;
}
.nav-header {
  width: 100%;
  height: 80px;
  line-height: 80px;
  background-image: url("/assets/images/header-ng.png");
  background-size: 100%;
  background-repeat: no-repeat;
  font-size: 0;
  color: #fff;
  text-align: center;
}
.nav-header .logo {
  display: inline-block;
  cursor: pointer;
}
.nav-chainType {
  position: absolute;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
}
.nav-menu {
  float: right;
}
.nav-item {
  display: inline-block;
  font-size: 14px;
  width: 120px;
  cursor: pointer;
}
/*.nav-item:hover{*/
/*border-bottom: 2px solid red;*/
/*}*/
.nav-item .icon-up {
  display: inline !important;
}
.nav-item .icon-down {
  display: none !important;
}
.nav-item:hover .icon-up {
  display: none !important;
}
.nav-item:hover .icon-down {
  display: inline !important;
}
.nav-menu-item {
  position: relative;
}
.nav-menu-item-option {
  display: none;
  position: absolute;
  width: 100%;
  top: 80px;
  left: 0;
  padding: 10px 0;
  font-size: 14px;
  color: #333;
  background-color: #fff;
  cursor: pointer;
  line-height: 32px;
  z-index: 9999999;
  box-sizing: border-box;
  margin: 0;
  list-style: none;
}
.nav-menu-item:hover .nav-menu-item-option {
  display: block;
}
.nav-menu-item-option option {
  cursor: pointer;
  height: 32px;
}
.nav-menu-item-option li:hover {
  color: #3498db !important;
}
.nav-item input {
  display: inline-block;
  width: 120px;
  padding-right: 10px;
  background-color: rgba(0, 0, 0, 0);
  color: #fff;
  text-align: center;
  border: none;
  outline: none;
  font-size: 14px;
}
.nav-item .select {
  position: absolute;
  display: none;
  top: 100%;
  left: 0;
  width: 100%;
  margin: 0;
  background-color: #fff;
  /*padding: 20px 0;*/
  z-index: 999999;
  padding: 10px 0;
  font-size: 14px;
  list-style: none;
  text-align: center;
}
.nav-item .select1 {
  position: absolute;
  display: none;
  top: 100%;
  left: 50%;
  width: 100%;
  margin: 0;
  background-color: #fff;
  /*padding: 20px 0;*/
  z-index: 999999;
  padding: 10px 0;
  font-size: 14px;
  list-style: none;
  text-align: center;
}
.nav-item .select li {
  height: 32px;
  line-height: 32px;
  background-color: #fff;
  /*padding: 20px 0;*/
  font-size: 14px;
  color: #333;
}
.nav-item .select1 li {
  height: 32px;
  line-height: 32px;
  background-color: #fff;
  /*padding: 20px 0;*/
  font-size: 0;
  color: #333;
}

.nav-item .select li:hover {
  color: #3498db !important;
}

.nav-menu-item:hover .select {
  display: block;
}
.nav-item:hover .select1 {
  display: block;
}
.nav-item .name:hover {
  color: #3498db !important;
}
.nav-item .delete:hover {
  color: #3498db !important;
}
.nav-item .name {
  display: inline-block;
  width: 50%;
  cursor: pointer;
  font-size: 14px;
}
.nav-item .delete {
  display: inline-block;
  width: 50%;
  cursor: pointer;
  font-size: 14px;
}
.nav-item .add {
  font-size: 14px !important;
  cursor: pointer;
}
.nav-item .add:hover {
  color: #3498db !important;
}

.icon {
  position: absolute;
  left: 101%;
  top: 42%;
}
</style>
