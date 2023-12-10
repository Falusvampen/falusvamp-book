<template>
  <div class="sidebar" :class="{ active: isActive }">
    <!-- Sidebar content -->
    <ul class="friends-list">
      <!-- Friend list items -->
      <li>Friend 1</li>
      <li>Friend 2</li>
      <li>Friend 3</li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isActive: false, // Flag to control the visibility of the sidebar
    };
  },
};
</script>

<style>
.sidebar {
  position: fixed;
  top: 56px; /* Adjust the value according to your navbar height */
  right: 0;
  width: 300px;
  height: 100vh;
  background-color: #f0f0f0;
  transition: transform 0.3s ease-out;
  transform: translateX(100%);
  z-index: 999;
}

.sidebar.active {
  transform: translateX(0);
}

.friends-list {
  list-style-type: none;
  padding: 0;
}

.friends-list li {
  padding: 10px;
  border-bottom: 1px solid #ccc;
}

.friends-list li:last-child {
  border-bottom: none;
}
</style>
