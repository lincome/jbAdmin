import { defineStore } from 'pinia'
import router from '@/router'
import i18n from '@/i18n'

const elementPlusLoacleList = await batchImport(import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs'))
export const useLanguageStore = defineStore('language', {
  state: () => {
    return {
      language: getLanguage(),
      //elementPlusLoacleList: import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs'),
      elementPlusLoacleList: elementPlusLoacleList,
    }
  },
  getters: {
    // elementPlusLocale: async (state) => {
    //   switch (state.language) {
    //     default:
    //       //return (await import(/* @vite-ignore */'../../node_modules/element-plus/dist/locale/' + state.language + '.mjs')).default
    //       return (await import(/* @vite-ignore */'/node_modules/element-plus/dist/locale/' + state.language + '.mjs')).default
    //   }
    // },
    elementPlusLocale: (state) => {
      switch (state.language) {
        default:
          //console.log(state.elementPlusLoacleList)
          //return (<any>state.elementPlusLoacleList)['/node_modules/element-plus/dist/locale/' + state.language + '.min.mjs'].default
          return state.elementPlusLoacleList[state.language]
      }
    }
  },
  actions: {
    //改变语言
    changeLanguage(language: string) {
      if (getLanguage() == language) {
        return
      }
      setLanguage(language)
      this.language = language
      //i18n.global.locale = language //当i18n设置legacy: false，要使用i18n.global.locale.value
      i18n.global.locale.value = language

      document.title = this.getWebTitle(router.currentRoute.value.fullPath)
      /**
       * 下面这几种情况，需要使用router.go(0)，强制刷新页面
       *    路由设置标题时，不能动态刷新
       *    部分接口，不能动态刷新
       */
      //router.go(0)
    },
    //获取页面标题
    getMenuTitle(menu: any) {
      if (menu) {
        return menu?.title?.[i18n.global.locale.value] ?? menu.menuName
      }
      return ''
    },
    //获取页面标题
    getPageTitle(fullPath: string) {
      const menu = useAdminStore().menuList.find((item) => {
        return item.url == fullPath
      }) ?? (<any>router).getRoutes().find((item: any) => {
        return fullPath.indexOf(item.path) === 0
      })?.meta?.menu
      return this.getMenuTitle(menu)
    },
    //获取网站标题
    getWebTitle(fullPath: string) {
      let webTitle = i18n.global.t('config.webTitle')
      const title = this.getPageTitle(fullPath)
      if (title) {
        webTitle += '-' + title
      }
      return webTitle
    },
  },
})
