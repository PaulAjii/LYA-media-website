<template>
    <div class="absolute h-screen top-0 z-100 bg-white py-2 transition-all duration-500 ease-in-out md:hidden"
        :class="isOpen ? 'left-0 w-screen' : '-left-full w-0'">
        <div class="flex items-center justify-between mb-14 px-6">
            <LogoComponent :isOpen="isOpen" />
            <PhX class="cursor-pointer font-bold text-gray-700 hover:text-red-700" size="30" weight="bold"
                @click.prevent="closeSideMenu" />
        </div>

        <ul class="nav__links gap-10 flex flex-col items-start font-oswald text-4xl font-bold mb-10 px-8">
            <li>
                <NuxtLink to="/" @click="closeSideMenu"
                    :class="route.path === '/' ? 'text-blue-700' : 'text-gray-600 hover:text-gray-900'">
                    ABOUT
                </NuxtLink>
            </li>
            <li>
                <NuxtLink to="/teachings" @click="closeSideMenu"
                    :class="route.path === '/teachings' ? 'text-blue-700' : 'text-gray-600 hover:text-gray-900'">
                    TEACHINGS
                </NuxtLink>
            </li>
            <li>
                <NuxtLink to="/choir-ministrations" @click="closeSideMenu"
                    :class="route.path === '/choir-ministrations' ? 'text-blue-700' : 'text-gray-600 hover:text-gray-900'">
                    CHOIR-MINISTRATIONS
                </NuxtLink>
            </li>
        </ul>

        <div class="px-8">
            <p class="mb-2 text-xs">SOCIALS:</p>
            <div class="flex items-center gap-6 font-oswald">
                <NuxtLink to="https://www.facebook.com/OkeAlafiaYouthChurchIB" rel="noopener" target="_blank" external
                    title="Facebook">
                    <PhFacebookLogo size="30" class="text-gray-500 hover:text-gray-900" weight="bold" />
                </NuxtLink>
                <NuxtLink to="https://www.instagram.com/loyolayouth_assembly" rel="noopener" target="_blank" external
                    title="Instagram">
                    <PhInstagramLogo size="30" class="text-gray-500 hover:text-gray-900" weight="bold" />
                </NuxtLink>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { PhFacebookLogo, PhInstagramLogo, PhX } from '@phosphor-icons/vue';

const props = defineProps<{
    isOpen: boolean
}>()
const route = useRoute()
const emit = defineEmits<{
    update: [isOpen: boolean]
}>()

const closeSideMenu = () => {
    emit("update", false)
}
</script>