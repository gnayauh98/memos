import { defineConfig, presetUno, presetAttributify, presetIcons } from 'unocss'
import transformerVariantGroup from '@unocss/transformer-variant-group'

export default defineConfig({
    presets: [
        presetAttributify(),
        presetUno(),
        presetIcons()
    ],
    transformers: [
        transformerVariantGroup()
    ],
    safelist: ["i-lucide:clipboard", "i-lucide:link", "i-lucide:x"]
})