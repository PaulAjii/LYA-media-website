<template>
    <p class="bg-gray-100 px-6 py-4 text-sm text-gray-400 font-logo">
        Copyright &copy; {{ date }}. Loyola Youth Assembly
    </p>
</template>

<script setup lang="ts">
const newDate = new Date()
const date = newDate.getFullYear()

</script>