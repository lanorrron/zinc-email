<template>
  <button
    v-bind="rest"
    :class="[ theme.colors.button, theme.colors.buttonHover,
      'rounded-md p-2  text-white',
      className,
    ]"
  >
    {{ label }}
  </button>
</template>

<script setup lang="ts">
import { theme } from '@/styles/theme';
import type { ButtonHTMLAttributes } from 'vue';

const { label, className, ...rest } = defineProps<
  {
    label: string;
    className?: string; 
  } & /* @vue-ignore */ ButtonHTMLAttributes
>();
</script>
