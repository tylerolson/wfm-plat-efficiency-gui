<script setup lang="ts">
import { FetchVendor, GetVendor } from "../../wailsjs/go/main/App"

import { onMounted, ref } from "vue"
import { wfmplatefficiency } from "../../wailsjs/go/models"

const props = defineProps<{
  name: string
}>()

const items = ref<wfmplatefficiency.Item[]>([])

function getItems() {
  GetVendor(props.name).then((result: wfmplatefficiency.Vendor) => {
    items.value = result.items
  })
}

async function fetch() {
  await FetchVendor(props.name).then(result => {})
  getItems()
}

onMounted(() => {
  getItems()
})
</script>

<template>
  <div class="flex justify-center">
    <div class="w-full overflow-scroll">
      <p>{{ name }}</p>
      <button @click="fetch()" class="cursor-pointer rounded-md bg-gray-400 px-4 py-1">Fetch</button>

      <table class="w-full">
        <thead class="bg-gray-700 text-xs text-gray-400 uppercase">
          <tr>
            <th class="py-3">Name</th>
            <th class="py-3">Type</th>
            <th class="py-3">Standing</th>
            <th class="py-3">Average Volume</th>
            <th class="py-3">Average Price</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.name">
            <td>{{ item.name }}</td>
            <td>{{ item.type }}</td>
            <td>{{ item.standing }}</td>
            <td>{{ item.AvgVol }}</td>
            <td>{{ item.WeightedAvgPrice.toFixed(3) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
