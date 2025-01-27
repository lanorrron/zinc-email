<template>
    <input
      type="text"
      :value="value"
      :placeholder="placeHolder"
      @input="handleInput"
      :class="[theme.colors.border, theme.colors.card]"
      class="ounded-md text-white focus:outline-none border-opacity-30 p-2"
    />
  </template>
  
  <script setup lang="ts">
import { theme } from '@/styles/theme';

  const props = defineProps<{
    value: string;
    placeHolder: string;
  }>();

  const emit = defineEmits<{
    (event: 'update:value', value: string): void;
  }>();
  
  const handleInput = (event: Event) => {
    const newValue = (event.target as HTMLInputElement).value;
    emit('update:value', newValue);
  };
  </script>
  