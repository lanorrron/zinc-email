<template>
    <div v-if="emails.length === 0" class="flex items center justify-center p-2">
        <p class="text-white">!There are no emails available at this time!</p>
    </div>
    <div v-else class="grid grid-cols-1 gap-2">
        <div v-for="(itemEmail, index) in emailsToShow" :key="index">
            <Card class="text-white flex flex-wrap">
                <div class="grid grid-cols-1 gap-2 w-full">
                    <p class="text-white text-xl font-mono" v-html="highlightText( itemEmail.subject, valueToSearched)"></p>
                     <div v-if="isTruncate[index]">
                        <p v-html="highlightText(itemEmail.body, valueToSearched)"
                        v-bind:class="{ 'line-clamp-3': isTruncate[index], 'line-clamp-none': !isTruncate[index] }"
                        class="font-sans"></p>
                     </div>
                     <div v-else>
                        <p v-for="(value, key) in itemEmail" :key="key">
                            <strong>{{ key }}:</strong> <span v-html="highlightText(value, valueToSearched)"></span>
                        </p>
                     </div>
                  
                    <div class="flex items-center gap-2 justify-between">
                        <div class="flex gap-2">
                            <div class="flex items-center gap-2">
                                <i class="fa fa-user-o text-red-900" aria-hidden="true"></i>
                                <p v-html="highlightText(formatUserName(itemEmail.from), valueToSearched)"></p>
                            </div>
                            <div class="flex items-center gap-2">
                                <i class="fa fa-calendar text-red-900" aria-hidden="true"></i>
                                <p v-html="highlightText(itemEmail.date, valueToSearched)"></p>
                            </div>
                        </div>
                        <div>
                            <Button :label="isTruncate[index] ? '+' : '-'" @click="togleTruncate(index)"
                                class-name="px-1 py-0"></Button>
                        </div>
                    </div>
                </div>
            </Card>
        </div>
        <Pagination :currentPage= "currentPage" v-bind:totalPages="totalPages" @update:currentPage="handleNavigation"></Pagination>
    </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import type { EmailType } from '../types/Email.type';
import Card from '@/components/Card.vue';
import Button from '@/components/Button.vue';
import Pagination from '@/components/Pagination.vue';

const props = defineProps<{
    emails: EmailType[];
    valueToSearched?: string;
}>();

const isTruncate = ref<boolean[]>([]);
const currentPage = ref(1)
const emailsPerPage = 5

const emailsToShow = computed(() => {
  const start = (currentPage.value - 1) * emailsPerPage;
  const end = start + emailsPerPage;
  return props.emails.slice(start, end);
});

const handleNavigation =(newPage: number)=>{
    currentPage.value = newPage
}

const totalPages = computed(() => {
  return Math.ceil(props.emails.length / emailsPerPage);
});

watch(() => props.emails, (newEmails: EmailType[]) => {
    isTruncate.value = newEmails.map(() => true);
}, { immediate: true });

const togleTruncate = (index: number) => {
    isTruncate.value[index] = !isTruncate.value[index];
};

const highlightText = (emailOrText: EmailType | string, search: string | undefined): EmailType | string => {
    if (!search || !emailOrText) return emailOrText;

    const words = search.trim().split(/\s+/).filter(Boolean);
    if (words.length === 0) return emailOrText;

    const regex = new RegExp(`(${words.join('|')})`, 'gi');

    if (typeof emailOrText === 'string') {
        return emailOrText.replace(regex, `<span class="bg-purple-500 rounded-sm text-purple-500 font-bold" style="--tw-bg-opacity:0.12">$1</span>`);
    }

    const highlightedEmail = Object.fromEntries(
        Object.entries(emailOrText).map(([key, value]) =>
            typeof value === 'string'
                ? [key, value.replace(regex, `<span class="bg-purple-500 rounded-sm text-purple-500 font-bold" style="--tw-bg-opacity:0.12">$1</span>`)]
                : [key, value]
        )
    );

    return highlightedEmail as EmailType;
};
const formatUserName = (value:string)=>{
    const userName = value.split("@")[0].replace(/\./g, " ");
    return userName;
}
</script>
