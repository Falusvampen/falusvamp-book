<template>
  <div class="chat-bubble">
    <div class="profile">
      <img :src="user.avatar" :alt="user.username" class="avatar" />
      <span class="username">{{ user.username }}</span>
    </div>
    <div class="chat-content">
      <div v-for="(message, index) in messages" :key="index" class="message">
        <span class="message-text">{{ message.text }}</span>
        <span class="message-time">{{ message.time }}</span>
      </div>
    </div>
    <div class="chat-input">
      <input
        v-model="newMessage"
        @keydown.enter="sendMessage"
        placeholder="Type your message..."
      />
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    user: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      messages: [],
      newMessage: "",
    };
  },
  methods: {
    sendMessage() {
      if (this.newMessage.trim() !== "") {
        const currentTime = new Date().toLocaleTimeString([], {
          hour: "2-digit",
          minute: "2-digit",
        });
        this.messages.push({ text: this.newMessage, time: currentTime });
        this.newMessage = "";
      }
    },
  },
};
</script>

<style scoped>
.chat-bubble {
  display: flex;
  flex-direction: column;
  width: 300px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 16px;
}

.profile {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.avatar {
  width: 32px;
  height: 32px;
  object-fit: cover;
  border-radius: 50%;
  margin-right: 8px;
}

.username {
  font-weight: bold;
  font-size: 14px;
}

.chat-content {
  flex: 1;
  overflow-y: auto;
  margin-bottom: 12px;
}

.message {
  display: flex;
  flex-direction: column;
  margin-bottom: 8px;
}

.message-text {
  font-size: 14px;
}

.message-time {
  font-size: 12px;
  color: #999;
}

.chat-input {
  display: flex;
}

input {
  flex: 1;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 8px;
  font-size: 14px;
}

button {
  margin-left: 8px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 8px 16px;
  cursor: pointer;
  font-size: 14px;
}
</style>
