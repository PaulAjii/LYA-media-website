<template>
    <header :class="[
        'w-full h-20 px-6 py-2 flex items-center justify-between fixed z-100 transition-all duration-300',
        !isTransparent
            ? 'bg-white border-b border-gray-200 shadow-sm'
            : 'backdrop-blur-md bg-white/10 border-b border-white/10'
    ]">
        <LogoComponent :isOpen="isOpen" />
        <!-- Desktop Mode Menu -->
        <div class="items-center justify-between hidden md:flex md:w-3/5">
            <div class="nav__links gap-3 flex items-center font-logo text-sm font-semibold">
                <NuxtLink to="/" :class="[
                    'transition-colors duration-300',
                    !isTransparent ? 'text-gray-700 hover:text-black' : 'text-white/80 hover:text-white/80'
                ]">
                    ABOUT
                </NuxtLink>
                <NuxtLink to="/teachings" :class="[
                    'transition-colors duration-300',
                    !isTransparent ? 'text-gray-700 hover:text-black' : 'text-white/80 hover:text-white/80'
                ]">
                    TEACHINGS
                </NuxtLink>
                <NuxtLink to="/choir-ministrations" :class="[
                    'transition-colors duration-300',
                    !isTransparent ? 'text-gray-700 hover:text-black' : 'text-white/80 hover:text-white/80'
                ]">
                    CHOIR-MINISTRATIONS
                </NuxtLink>
            </div>
            <div class="flex items-center gap-4">
                <NuxtLink to="https://www.facebook.com/OkeAlafiaYouthChurchIB" rel="noopener" target="_blank" external>
                    <PhFacebookLogo size="25" :class="[
                        'transition-colors duration-300',
                        !isTransparent ? 'text-gray-600 hover:text-black' : 'text-white/80 hover:text-white',
                    ]" weight="bold" />
                </NuxtLink>
                <NuxtLink to="https://www.instagram.com/loyolayouth_assembly" rel="noopener" target="_blank" external>
                    <PhInstagramLogo size="25" :class="[
                        'transition-colors duration-300',
                        !isTransparent ? 'text-gray-600 hover:text-black' : 'text-white/80 hover:text-white',
                    ]" weight="bold" />
                </NuxtLink>
            </div>
        </div>

        <!-- Mobile Menu -->
        <SideMenu :isOpen="isOpen" @update="handleUpdate" />

        <button type="button" class="cursor-pointer transition-colors duration-300 md:hidden"
            @click.prevent="() => isOpen = true">
            <PhList size="25"
                :class="!isTransparent ? 'text-gray-700 hover:text-black' : 'text-white/80 hover:text-white'"
                weight="bold" />
        </button>
    </header>
</template>

<script setup lang="ts">
import { PhFacebookLogo, PhInstagramLogo, PhList } from '@phosphor-icons/vue';

const { isScrolled } = useScrollHeader()
const { isDark } = useHeaderTheme()
const isOpen = ref(false)

const handleUpdate = (newValue: boolean) => {
    isOpen.value = newValue
}

const isTransparent = computed(() => isDark.value && !isScrolled.value)
</script>