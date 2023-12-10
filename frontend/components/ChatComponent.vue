<template>
  <div class="chat-component">
    <div class="chat-header">
      <h2 class="chat-title">Chat</h2>
      <button class="chat-close" @click="toggleChat">
        <iconsClose />
      </button>
    </div>
    <div class="chat-content"></div>
  </div>
</template>

<script>
export default {
  methods: {
    toggleChat() {
      this.$emit("toggle");
    },
  },
};
</script>

<style scoped>
.chat-component {
  position: fixed;
  top: 0;
  right: 0;
  width: 320px;
  height: 100vh;
  background-color: #f0f0f0;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  z-index: 100;
  overflow-y: auto;
}

.chat-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px;
  background-color: #4267b2;
  color: #fff;
}

.chat-title {
  margin: 0;
  font-size: 18px;
}

.chat-close {
  border: none;
  background-color: transparent;
  color: #fff;
  cursor: pointer;
  font-size: 20px;
}

.chat-content {
  padding: 10px;
}
</style>
