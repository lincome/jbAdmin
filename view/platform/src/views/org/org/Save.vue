<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data,
    } as { [propName: string]: any },
    rules: {
        org_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 60, message: t('validation.max.string', { max: 60 }) },
        ],
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            try {
                if (param?.id) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/org/org/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/org/org/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } finally {
                saveForm.loading = false
            }
        })
    },
})

const saveDrawer = reactive({
    ref: null as any,
    size: useSettingStore().saveDrawer.size,
    beforeClose: (done: Function) => {
        if (useSettingStore().saveDrawer.isTipClose) {
            ElMessageBox.confirm('', {
                type: 'info',
                title: t('common.tip.configExit'),
                center: true,
                showClose: false,
            }).then(() => done())
        } else {
            done()
        }
    },
    buttonClose: () => saveDrawer.ref.handleClose(), //saveCommon.visible = false //不会触发beforeClose
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('org.org.name.org_name')" prop="org_name">
                    <el-input v-model="saveForm.data.org_name" :placeholder="t('org.org.name.org_name')" maxlength="60" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('org.org.name.is_stop')" prop="is_stop">
                    <el-switch
                        v-model="saveForm.data.is_stop"
                        :active-value="(tm('common.status.whether') as any[])[1].value"
                        :inactive-value="(tm('common.status.whether') as any[])[0].value"
                        :active-text="(tm('common.status.whether') as any[])[1].label"
                        :inactive-text="(tm('common.status.whether') as any[])[0].label"
                        :inline-prompt="true"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">{{ t('common.save') }}</el-button>
        </template>
    </el-drawer>
</template>
