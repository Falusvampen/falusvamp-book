<template>
  <div class="modal" v-if="showModal">
    <div class="modal-content">
      <span class="close" @click="closeModal">&times;</span>

      <div class="post">
        <div class="post-meta">
          <profileImage />
          <p>Posted by: {{ user }}</p>
          <p>Created at: {{ createdAt }}</p>
        </div>
        <h2 class="post-title">{{ title }}</h2>
        <p class="post-content">
          <img :src="image" class="post-image" style="width: 100%" />
          {{ content }}
        </p>

        <div class="post-buttons">
          <button @click="incrementLikes" class="post-like-button">
            <iconsThumbsUp />
            <span class="post-like-count">{{ likes }}</span>
          </button>
          <button @click="incrementDislikes" class="post-dislike-button">
            <iconsThumbsDown />
            <span class="post-dislike-count">{{ dislikes }}</span>
          </button>
        </div>
      </div>

      <div class="comments-section">
        <h3>Comments</h3>
        <div class="comment" v-for="comment in comments" :key="comment.id">
          <p>
            <strong>{{ comment.user }}</strong> {{ comment.text }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    content: String,
    user: String,
    createdAt: String,
    image: String,
    comments: Array,
    showModal: Boolean,
  },
  data() {
    return {
      likes: 0,
      dislikes: 0,
    };
  },
  methods: {
    incrementLikes() {
      this.likes++;
    },
    incrementDislikes() {
      this.dislikes++;
    },
    closeModal() {
      this.$emit("close-modal");
    },
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");

.modal {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 9999;
}

.modal-content {
  background-color: #f5f5f5;
  border-radius: 4px;
  padding: 20px;
  width: 80%;
  max-width: 500px;
}

.close {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 20px;
  cursor: pointer;
}

.post {
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 4px;
  font-family: "Poppins", sans-serif;
}

.post-title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 10px;
}

.post-content {
  font-size: 16px;
  line-height: 1.5;
  margin-bottom: 20px;
}

.post-meta {
  color: #777;
  font-size: 14px;
  margin-bottom: 10px;
}

.post-buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.post-like-button,
.post-dislike-button {
  background-color: transparent;
  border: none;
  color: #888;
  font-size: 14px;
  cursor: pointer;
  display: flex;
  align-items: center;
}

.post-like-count,
.post-dislike-count {
  margin-left: 5px;
  font-weight: bold;
}

.comments-section {
  margin-top: 20px;
}

.comment {
  margin-bottom: 10px;
}
</style>
