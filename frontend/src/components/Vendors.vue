<script setup lang="ts">
import { GetVendorNames } from "../../wailsjs/go/main/App";

import { onMounted, ref } from "vue";

const vendors = ref<string[]>([]);
onMounted(() => {
  GetVendorNames().then((result) => {
    vendors.value = result;
  });

  console.log(`the component is now mounted.`);
});
</script>

<template>
  <div>
    <span>Vendors:</span>
    <li v-for="v in vendors">{{ v }}</li>
  </div>
</template>
