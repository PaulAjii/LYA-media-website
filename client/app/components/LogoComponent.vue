<template>
    <NuxtLink class="flex gap-2 items-center transition-all duration-300 md:w-1/2" to="/">
        <img src="../assets/images/logo_use.png" alt="Logo" :class="[
            'w-15 h-15 object-cover transition-all duration-300',
        ]" />
        <div v-if="isTransparent" class="flex flex-col gap-0.5 text-xs font-logo text-white/80 hover:text-white">
            <span>Loyola Youth Assembly (LYA)</span>
            <span>CAC Oke Alaafia, Loyola, Ibadan</span>
        </div>
        <div v-else-if="isOpen" class="flex flex-col gap-0.5 text-xs font-logo text-gray-700 hover:text-black">
            <span>Loyola Youth Assembly (LYA)</span>
            <span>CAC Oke Alaafia, Loyola, Ibadan</span>
        </div>
        <div v-else class="flex flex-col gap-0.5 text-xs font-logo text-gray-700 hover:text-black">
            <span>Loyola Youth Assembly (LYA)</span>
            <span>CAC Oke Alaafia, Loyola, Ibadan</span>
        </div>
    </NuxtLink>
</template>

<script setup lang="ts">
const { isScrolled } = useScrollHeader()
const { isDark } = useHeaderTheme()

const isTransparent = computed(() => isDark.value && !isScrolled.value)

const props = defineProps<{
    isOpen: boolean
}>()

</script>