import i18n from '@/i18n';
const apiList = await batchImport(import.meta.globEager('@/api/**/*.ts'))

/**
 * 请求接口
 * @param apiCode   接口标识。apiList内的键用'.'拼接组成。例如：login.getEncryptStr
 * @param data  请求参数
 * @param isErrorHandle 错误处理，默认true。有时调用接口的位置报错需要特殊处理，传false则会抛出错误，可在调用处捕获错误再处理。
 * @returns 
 */
export const request = async (apiCode: string, data?: {}, isErrorHandle: boolean = true) => {
    //const apiList = batchImport(import.meta.globEager('@/api/**/*.ts')) //放外面去。这样每次调用都不要重新加载了
    let apiMethod: any = apiList;
    for (const value of apiCode.split('.')) {
        if (!(value in apiMethod)) {
            break;
        }
        apiMethod = apiMethod[value]
    }

    if (typeof apiMethod !== 'function') {
        errorHandle(new Error(i18n.global.t('error.apiFunctionNoFind')))
        return false
    }

    try {
        return await apiMethod(data)
    } catch (error) {
        if (isErrorHandle) {
            errorHandle(<Error>error)
            //throw error
            return false
        } else {
            throw error
        }
    }
}

/*--------使用方式 开始--------*/
/* request('index.index', data)

request('index.index', undefined, false)
    .then((data) => {
        console.log(data)
    })
    .catch((error) => {
        errorHandle(<Error>error)
    })

try {
    await request('index.index', data, false)
} catch (error) {
    errorHandle(<Error>error)
} */
/*--------使用方式 结束--------*/