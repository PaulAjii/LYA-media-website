import { ref, onMounted, onUnmounted } from "vue"

export function useScrollHeader(threshold = 600) {
    const isScrolled = ref(false)

    function onScroll() {
        isScrolled.value = window.scrollY > threshold
    }

    onMounted(() => window.addEventListener("scroll", onScroll))
    onUnmounted(() => window.removeEventListener("scroll", onScroll))

    return { isScrolled }
}
