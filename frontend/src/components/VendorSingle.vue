<script setup lang="ts">
import { FetchVendor, GetVendor } from "../../wailsjs/go/main/App"

import { onMounted, ref, computed } from "vue"
import { wfmplatefficiency } from "../../wailsjs/go/models"
// import { LogPrint } from "../../wailsjs/runtime/"

const props = defineProps<{
  name: string
}>()

const fetchButton = ref("Fetch")

const sortKey = ref("name")
const sortAsc = ref(true)

const items = ref<wfmplatefficiency.Item[]>([])

const sortedItems = computed(() => {
  if (!sortKey.value) {
    return [...items.value]
  }

  const sortedArray = [...items.value].sort((a, b) => {
    const valueA = a[sortKey.value as keyof wfmplatefficiency.Item]
    const valueB = b[sortKey.value as keyof wfmplatefficiency.Item]

    if (valueA < valueB) {
      return sortAsc.value ? -1 : 1
    }
    if (valueA > valueB) {
      return sortAsc.value ? 1 : -1
    }
    return 0
  })
  return sortedArray
})

function sort(key: keyof wfmplatefficiency.Item) {
  if (sortKey.value === key) {
    sortAsc.value = !sortAsc.value
  } else {
    sortKey.value = key
    sortAsc.value = true
  }
}

function getItems() {
  GetVendor(props.name).then((result: wfmplatefficiency.Vendor) => {
    items.value = result.items
  })
}

async function fetch() {
  fetchButton.value = "Fetching..."
  await FetchVendor(props.name)
  getItems()
  fetchButton.value = "Fetched"
}

onMounted(() => {
  getItems()
})
</script>

<template>
  <div class="flex justify-center">
    <div class="w-full overflow-scroll">
      <div class="flex items-center justify-center py-5">
        <span class="text-4xl">{{ name }}</span>
        <button
          @click="fetch()"
          class="ml-5 cursor-pointer rounded-md bg-gray-400 px-4 py-1 transition hover:bg-gray-500"
        >
          {{ fetchButton }}
        </button>
      </div>
      <table class="w-full">
        <thead class="bg-gray-700 text-xs text-gray-400 uppercase">
          <tr>
            <th @click="sort('name')" class="cursor-pointer py-3 select-none">Name</th>
            <th @click="sort('type')" class="cursor-pointer py-3 select-none">Type</th>
            <th @click="sort('standing')" class="cursor-pointer py-3 select-none">Standing</th>
            <th @click="sort('AvgVol')" class="cursor-pointer py-3 select-none">Average Volume</th>
            <th @click="sort('WeightedAvgPrice')" class="cursor-pointer py-3 select-none">Average Price</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in sortedItems" :key="item.name">
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
