<template>
  <div class="relative">
    <button
      @click="isOpen = !isOpen"
      class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded"
    >
      Notifications
    </button>

    <div
      v-if="isOpen"
      class="absolute right-0 mt-2 bg-white shadow-lg rounded-lg p-4 w-64 z-10"
    >
      <div v-if="notifications.length === 0" class="text-center">
        No new notifications
      </div>

      <ul v-else>
        <li v-for="notification in notifications" :key="notification.id">
          {{ notification.message }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
const isOpen = ref(false);
const notifications = ref([
  { id: 1, message: "Notification 1" },
  { id: 2, message: "Notification 2" },
  { id: 3, message: "Notification 3" },
]);
</script>

<style scoped>
.notification-dropdown {
  position: relative;
}
</style>
