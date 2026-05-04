export function useHeaderTheme() {
  const isDark = useState<boolean>('header-dark', () => false)

  function setDark() { 
    if (import.meta.client) {
      isDark.value = true 
    }
  }
  function setLight() { 
    if (import.meta.client) {
      isDark.value = false
    }
  }

  return { isDark, setDark, setLight }
}