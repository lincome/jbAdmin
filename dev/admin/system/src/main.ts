import { createApp } from "vue";
import App from "./App.vue";

import "./assets/main.css";
import { config } from "@/basic/functions";

const app = createApp(App);

import router from './router';
app.use(router);

import { createPinia } from 'pinia';
app.use(createPinia());

import i18n from './i18n';
app.use(i18n);

app.mount('#app');

/*-------- 动态加载图标 开始 --------*/
//app.component('autoiconEpLollipop', autoiconEpLollipop)
import * as epIconList from '@element-plus/icons-vue'
for (let [key, component] of (<any>Object).entries(epIconList)) {
    app.component('Ep' + key, component)
    app.component('AutoiconEp' + key, component)    //兼容图标插件unplugin-icons，如插件以后支持动态加载<component :is="图标标识变量"/>，不用修改代码
}
/*-------- 动态加载图标 结束 --------*/