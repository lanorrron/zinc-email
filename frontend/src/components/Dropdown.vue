<template>
    <div ref="dropdown" class="relative">
      <div 
        @click="toggleDropdown"
        :class="[theme.colors.border, theme.colors.card]"
        class="min-w-40 px-4 py-2 text-white rounded-md cursor-pointer truncate flex justify-between items-center"
      >
        {{ selectedOption?.label || "Select a option" }} 
         <i :class="isOpen?'fa fa-chevron-up': 'fa fa-chevron-down'" class="text-purple-500 text-opacity-30"></i>

      </div>
  
      <ul
      :class="[theme.colors.border, theme.colors.card]"
        v-if="isOpen"
        class="absolute z-10  mt-1 rounded-md min-w-40 "
      >
        <li
          v-for="option in options"
          :key="option.value"
          @click="selectOption(option)"
          class="px-4 py-2 text-white hover:bg-purple-500 hover:text-purple-500 cursor-pointer "
          style="--tw-bg-opacity: 0.12"
        >
          {{ option.label }}
 
        </li>
      </ul>
    </div>
  </template>
  
  <script setup lang="ts">
  import { theme } from "@/styles/theme";
import { ref, onMounted, onUnmounted } from "vue";

  const props = defineProps<{
    value: string | number;
    options: { value: string | number; label: string }[];
  }>();
  
  const emit = defineEmits<{
    (event: "update:value", value: string | number): void;
  }>();
  

  const isOpen = ref(false);
  const selectedOption = ref(props.options.find(option => option.value === props.value));
  const dropdown = ref<HTMLDivElement | null>(null);
  

  const toggleDropdown = (event: MouseEvent) => {
    event.stopPropagation();
    isOpen.value = !isOpen.value;
  };
  

  const selectOption = (option: { value: string | number; label: string }) => {
    selectedOption.value = option;
    emit("update:value", option.value); 
    isOpen.value = false; 
  };
  

  const handleClickOutside = (event: MouseEvent) => {
    if (dropdown.value && !dropdown.value.contains(event.target as Node)) {
      isOpen.value = false; 
    }
  };

  onMounted(() => {
    document.addEventListener("click", handleClickOutside);
  });
  
  onUnmounted(() => {
    document.removeEventListener("click", handleClickOutside);
  });
  </script>
  