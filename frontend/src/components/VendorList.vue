<script setup lang="ts">
import { GetVendorNames } from "../../wailsjs/go/main/App"

import VendorSingle from "./VendorSingle.vue"

import { onMounted, ref } from "vue"

const vendorNames = ref<string[]>([])

onMounted(() => {
  GetVendorNames().then((result: string[]) => {
    vendorNames.value = result
  })
})
</script>

<template>
  <div>
    <VendorSingle v-for="v in vendorNames" v-bind:name="v" />
  </div>
</template>
