<template>
    <div v-bind="$attrs" :class="[theme.colors.card, theme.colors.border]" class="p-4">
        <slot></slot>
    </div>
</template>

<script setup lang="ts">
import { theme } from '@/styles/theme';

</script>