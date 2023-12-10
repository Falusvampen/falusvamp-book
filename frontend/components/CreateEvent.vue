<template>
  <div class="post-creation">
    <h2>Create Event</h2>
    <div class="input-container">
      <label for="event-title">Title:</label>
      <input id="event-title" v-model="event.title" type="text" required />
    </div>
    <div class="input-container">
      <label for="event-description">Description:</label>
      <textarea
        id="event-description"
        v-model="event.description"
        required
      ></textarea>
    </div>
    <div class="input-container">
      <label for="event-dateTime">Day/Time:</label>
      <input
        type="datetime-local"
        id="event-dateTime"
        v-model="event.dateTime"
        required
      />
    </div>
    <button @click="createEvent">Create</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      event: {
        title: "",
        description: "",
        dateTime: "",
      },
    };
  },
  methods: {
    createEvent() {
      // Perform event creation logic here
      console.log("Creating event:", this.event);
      // Reset the form
      this.event.title = "";
      this.event.description = "";
      this.event.dateTime = "";
    },
  },
};
</script>

<style scoped>
.post-creation {
  margin-bottom: 20px;
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 4px;
  background-color: #f5f5f5;
  font-family: "Poppins", sans-serif;
  width: 400px;
  height: 500px;
}

.input-container {
  margin-bottom: 10px;
}

.input-container label {
  display: block;
  font-weight: bold;
}

input[type="text"],
textarea,
input[type="datetime-local"],
button {
  width: 100%;
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
  resize: none;
}

button {
  background-color: #007bff;
  color: white;
  padding: 8px 16px;
  border: none;
  cursor: pointer;
}

button:hover {
  background-color: #0069d9;
}
</style>
