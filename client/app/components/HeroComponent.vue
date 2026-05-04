<template>
    <div class="relative w-full h-svh overflow-hidden pt-20 -mt-12">
        <!-- Background image -->
        <img src="../assets/images/hero_banner.jpg" alt="Worship service" class="w-full h-full object-cover" />

        <!-- Overlay -->
        <div class="absolute inset-0 bg-linear-to-t from-black/95 via-black/40 to-black/70" />

        <!-- Content -->
        <div class="absolute bottom-1/2 right-1/2 translate-1/2 px-8 py-8 flex items-center flex-col">
            <span
                class="inline-block text-[10px] tracking-widest uppercase text-center text-white bg-gray-500 rounded-full px-3 py-1 mb-3 font-logo sm:text-xs">
                LOYOLA YOUTH ASSEMBLY
            </span>
            <p
                class="w-md uppercase text-gray-300 font-black font-oswald text-5xl flex items-center justify-center flex-col gap-2 mt-2 sm:text:6xl sm:w-l md:w-3xl md:text-8xl">
                <span class="text-blue-700">caring</span>
                <span>for</span>
                <span class="text-blue-700">the whole man</span>
            </p>
        </div>
    </div>
</template>