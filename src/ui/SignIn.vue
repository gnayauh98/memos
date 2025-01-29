<script lang="ts" setup>
import { reactive } from 'vue'
import Api, { RequestCode } from '../api'
import { useRouter } from 'vue-router'
import { useLocalStorage } from '@vueuse/core'

const router = useRouter()
const accessToken = useLocalStorage("access-token", () => {
    return localStorage.getItem("access-token")
})
const userInfo = useLocalStorage("user-info", {})

const form = reactive({
    email: {
        err: false,
        msg: "",
        value: ""
    },
    password: {
        err: false,
        msg: "",
        value: ""
    }
})

const isValid = (): boolean => {
    let valid = 0
    if (form.email.value.length === 0) {
        form.email.err = true
        form.email.msg = "用户邮箱地址不能为空"
        valid -= 1
    } else {
        form.email.err = false
        form.email.msg = ""
        valid += 1
    }

    if (form.password.value.length === 0) {
        form.password.err = true
        form.password.msg = "用户密码不能为空"
        valid -= 1
    } else {
        form.password.err = false
        form.password.msg = ""
        valid += 1
    }

    return valid == 2
}

const onSubmit = async (event: Event) => {
    event.preventDefault()
    // 验证数据
    if (!isValid()) {
        return
    }

    try {
        // 登录处理
        const { code, user, token } = await Api.user.signin(form.email.value, form.password.value)

        if (code !== RequestCode.REQUEST_SUCCESS) {
            // TODO 登录失败处理
            console.error("登录失败")
            return
        }
        // localStorage.setItem("access-token", token)
        accessToken.value = token
        userInfo.value = user

        // 登录成功跳转到首页
        router.replace({ name: 'home' })
    } catch {
        // TODO 登录异常处理
        console.error("登录异常");
    }
}

</script>

<template>
    <div class="max-w-768px mx-auto">
        <form @submit="onSubmit" class="flex flex-col gap-16px mx-auto max-w-300px bg-#fafafa p-32px rounded-4px">
            <div class="flex justify-center mt-32px mb-16px">
                <span>AI Memos</span>
            </div>
            <div>
                <input v-model="form.email.value" placeholder="邮箱地址"
                    class="w-full outline-none py-4px px-8px rounded-4px border-(1px transparent) focus-visible:border-(1px #cacaca)"
                    :class="{ 'border-(1px red)!': form.email.err }" />
                <div v-if="form.email.err" class="text-red text-0.75em">{{ form.email.msg }}</div>
            </div>
            <div>
                <input v-model="form.password.value" placeholder="密码"
                    class="w-full outline-none py-4px px-8px rounded-4px border-(1px transparent) focus-visible:border-(1px #cacaca)"
                    :class="{ 'border-(1px red)!': form.password.err }" />
                <div v-if="form.password.err" class="text-red text-0.75em">{{ form.password.msg }}</div>
            </div>
            <div class="flex text-12px text-blue">
                <span class="ml-auto cursor-pointer">忘记密码?</span>
            </div>
            <div>
                <button type="submit"
                    class="w-full outline-none py-4px px-8px rounded-4px text-white bg-green">登录</button>
            </div>
        </form>
    </div>
</template>