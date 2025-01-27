<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Search from './modules/email/ui/Search.vue';

const isScrolled = ref(false);

const handleScroll = () => {
  if (window.scrollY > 5) {
    isScrolled.value = true
  } else {
    isScrolled.value = false
  }
}
onMounted(()=>{
  addEventListener("scroll", handleScroll)
})
</script>

<template>
  <header class="fixed top-5 left-0 w-full z-10 ">
    <h1      :class="[
        'flex items-center justify-center text-5xl font-mono m-5 truncate', 
        isScrolled ? 'text-white drop-shadow-[0_0_10px_rgba(255,0,255,1)]' : 'text-white'
      ]">Email Server</h1>
  </header>

  <main class="flex items-center justify-center p-2 mt-28">
    <Search></Search>
  </main>
</template>
