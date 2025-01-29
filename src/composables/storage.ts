import { onMounted, onUnmounted, reactive } from 'vue'

export function useLocalStorage<T extends Object>(key: string, defaultValue: T | undefined, parser?: (value: string | null) => T) {

    parser ||= (value) => {
        if (value) {
            return JSON.parse(value)
        }
        // 如果没有设置值，返回默认值
        return defaultValue
    }

    const value = reactive(parser(localStorage.getItem(key)))

    function updateStorage(event: StorageEvent) {
        console.log(event)
    }

    function set(value: T | undefined) {
        if (value === undefined) {
            localStorage.removeItem(key)
        } else {
            localStorage.setItem(key, JSON.stringify(value))
        }
    }

    function get(): T {
        return parser?.(localStorage.getItem(key))!
    }

    onMounted(() => {
        console.log(121)
        window.addEventListener('storage', updateStorage)
    })

    onUnmounted(() => {
        console.log(123)
        window.removeEventListener('storage', updateStorage)
    })

    return { value, get, set }
}