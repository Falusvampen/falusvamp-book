<template>
  <div class="post-creation">
    <h2>Create Post</h2>
    <div class="input-container">
      <label for="post-title">Title:</label>
      <input id="post-title" v-model="title" type="text" />
    </div>
    <div class="input-container">
      <label for="post-content">Content:</label>
      <textarea id="post-content" v-model="content"></textarea>
    </div>
    <div class="input-container">
      <label for="post-image">Image:</label>
      <input
        id="post-image"
        ref="imageInput"
        type="file"
        @change="handleImageUpload"
        accept="image/*"
      />
    </div>
    <div class="input-container, radio-privacy">
      <label>Privacy:</label>
      <input type="radio" id="private" value="private" v-model="privacy" />
      <label for="private">Private</label>
      <input type="radio" id="public" value="public" v-model="privacy" />
      <label for="public">Public</label>
      <input
        type="radio"
        id="almost-private"
        value="almost private"
        v-model="privacy"
      />
      <label for="almost-private">Almost Private</label>
    </div>
    <button @click="createPost">Create</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: "",
      content: "",
      privacy: "private",
      picture: "",
    };
  },
  methods: {
    async createPost() {
      const postData = {
        title: this.title,
        content: this.content,
        privacy: this.privacy,
        picture: this.picture,
      };
      try {
        const response = await $fetch("http://localhost:8080/api/post/submit", {
          method: "POST",
          credentials: "include",
          body: JSON.stringify(postData),
          headers: {
            "Content-Type": "application/json",
          },
        });
        console.log(response);
      } catch (error) {
        console.log(error);
      }
    },

    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        if (!["image/jpeg", "image/png", "image/gif"].includes(file.type)) {
          alert("Invalid file type. Please select a JPEG, PNG, or GIF image.");
          return;
        }
        if (file.size > 500 * 1024) {
          alert("File is too large. Please select a file less than 500KB.");
          return;
        }
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => {
          const img = new Image();
          img.src = reader.result;
          img.onload = () => {
            const canvas = document.createElement("canvas");
            const ctx = canvas.getContext("2d");
            const MAX_WIDTH = 500;
            const MAX_HEIGHT = 500;
            let width = img.width;
            let height = img.height;
            if (width > height) {
              if (width > MAX_WIDTH) {
                height *= MAX_WIDTH / width;
                width = MAX_WIDTH;
              }
            } else {
              if (height > MAX_HEIGHT) {
                width *= MAX_HEIGHT / height;
                height = MAX_HEIGHT;
              }
            }
            canvas.width = width;
            canvas.height = height;
            ctx.drawImage(img, 0, 0, width, height);
            const dataUrl = canvas.toDataURL(file.type, 0.8);
            const base64String = dataUrl.replace(
              /^data:image\/\w+;base64,/,
              ""
            );
            this.picture = base64String;
          };
        };
      }
    },
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");

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

.radio-privacy {
  display: flex;
  flex-direction: row;
  align-items: center;
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
input[type="file"],
button {
  width: 100%;
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
  resize: none;
}

input[type="checkbox"] {
  margin-right: 5px;
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
