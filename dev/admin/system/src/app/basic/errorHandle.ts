import { useUserStore } from '@/stores/user'

/**
 * 错误处理
 * @param {*} err 错误信息。格式：JSON；包含字段：{ code: 9999, msg: '失败', data: {} }
 */
export const errorHandle = async (err) => {
    try {
        const errMsg = JSON.parse(err.message)
        if (typeof errMsg.code !== 'undefined') {
            switch (errMsg.code) {
                case 4000:
                    /* ElMessageBox.alert(errMsg.msg, '确认登出', {
                        confirmButtonText: '重新登录',
                        type: 'warning'
                    }).then(async () => {
                        await store.dispatch('user/logout')
                    }).catch(async () => {
                        await store.dispatch('user/logout')
                    }) */
                    useUserStore().logout(getCurrentPath())
                    ElMessage.error(errMsg.msg)
                    break
                default:
                    ElMessage.error(errMsg.msg)
                    break
            }
        }
    } catch (e) {
        ElMessage.error(e.message);
    }
}
