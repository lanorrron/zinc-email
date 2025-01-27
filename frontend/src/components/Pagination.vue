<template>
    <div class="flex justify-center items-center gap-4" >
        <Button label="Previous" :disabled="currentPage === 1" :onclick="goToPrevious"/>
        <Card class="px-2 py-2">
            <p class="text-white" >
                {{ currentPage }} / {{ totalPages }}
            </p>
          </Card>
        <Button label="Next" :disabled="currentPage === totalPages" :onclick="goToNext"/>
    </div>
</template>
<script setup lang="ts">
import Button from './Button.vue';
import Card from './Card.vue';

const props = defineProps<{
    currentPage: number;
    totalPages: number;
}>();

const emit = defineEmits<{
  (event: 'update:currentPage', value: number): void;
}>();

const goToPrevious = () => {
  if (props.currentPage > 1) {
    emit('update:currentPage', props.currentPage - 1);
  }
};

const goToNext = () => {
  if (props.currentPage < props.totalPages) {
    emit('update:currentPage', props.currentPage + 1);
  }
}
</script>