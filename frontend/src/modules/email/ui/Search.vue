<template>
    <div class="flex flex-col sm:w-10/12 md:w-8/12 lg:w-5/12 gap-2 ">
        <div class="flex justify-between gap-x-4 gap-y-2 flex-col sm:flex-row">
        
            <div class="flex flex-col">
                <Dropdown v-model:value="valueSelected" :options="options"></Dropdown>
                <div v-if="errorMessageListIndex">
                <p class="text-red-500 text-center">{{ errorMessageListIndex }}</p>
            </div>
            </div>

            <div class="text-white font-mono text-sm truncate hidden sm:block ">
                <p>Total Index: {{ totalIndex }}</p>
                <p>Total Emails: {{ totalEmail }}</p>
            </div>
        </div>
        <div class="flex gap-2">
            <Search class="w-full" placeHolder="Search Email..." v-model:value="valueToSearch"></Search>
            <Button label="Search" @click="fetchSearchEmails({ nameIndex: valueSelected, query: valueToSearch})" />
        </div>
        <div>
            <div v-if="errorMessageSearch">
                <Card>
                    <p class="text-red-500 text-xl">
                        {{ errorMessageSearch }}
                    </p>
                </Card>
            </div>
            <div v-else>
                <ViewEmail :emails="emails" :value-to-searched="valueToSearch"></ViewEmail>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Dropdown from '@/components/Dropdown.vue';
import Search from '@/components/Input.vue';
import Button from '@/components/Button.vue';
import ViewEmail from './ViewEmail.vue';
import { onMounted, ref, watch } from 'vue';
import type { EmailType, SearchRequestType, SearchResponseType } from '../types/Email.type';
import { listIndex, searchEmail } from '../services/EmailService'

// list index
const valueSelected = ref<string>("")
const emails = ref<EmailType[]>([])
const options = ref<{ value: string, label: string }[]>([]);
const errorMessageListIndex = ref<string>("")
const errorMessageSearch = ref<string>("")

const fetchDropdownOptions = async () => {
    try {
        const result = await listIndex();
        options.value = result.map((item: string) => ({
            value: item,
            label: item
        }));
        totalIndex.value = options.value.length
    } catch (error: any) {
        errorMessageListIndex.value = error?.response?.data?.message || "failed list index"
    }
}
// Search Emails
const fetchSearchEmails = async (params: SearchRequestType) => {
    try {
        const result: SearchResponseType = await searchEmail(params)
        emails.value = result.hits.hits.map((item) => item._source)
        totalEmail.value = result.hits.total.value

    } catch (error: any) {
        errorMessageSearch.value = error?.response?.data?.message || "failed to search emails"
    }
}

onMounted(() => {
    fetchDropdownOptions();
});

watch(valueSelected, (newValue: string) => {
    if (newValue) {
        fetchSearchEmails({ nameIndex: newValue, query: valueToSearch.value})
    }
})

let totalIndex = ref<number>(0)
let totalEmail = ref<number>(0)
const valueToSearch = ref<string>("")

</script>