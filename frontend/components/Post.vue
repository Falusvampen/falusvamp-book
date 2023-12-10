<template>
  <div class="post" @mouseover="hover = true" @mouseleave="hover = false">
    <div class="post-meta">
      <div class="profile-and-button">
        <profileImage />
        <button class="follow-button">Follow</button>
      </div>
      <p>Posted by: {{ user }}</p>
      <p>Created at: {{ createdAt }}</p>
    </div>
    <h2 class="post-title" @click="showModal = true">{{ title }}</h2>
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

    <div class="create-comment">
      <textarea
        v-model="commentText"
        class="comment-textarea"
        placeholder="Write a comment..."
      ></textarea>
      <button @click="createComment" class="comment-button">Comment</button>
    </div>

    <button
      v-if="comments.length > 0"
      @click="showComments = !showComments"
      class="show-comments-button"
    >
      {{ showComments ? "Hide Comments" : "Show Comments" }}
    </button>
    <p v-else class="no-comments">No comments available</p>

    <div v-if="showComments" class="post-comments">
      <Comment
        v-for="comment in comments"
        :key="comment.id"
        :text="comment.text"
        :author="comment.author"
        :date="comment.date"
        :profileImage="comment.profileImage"
        :image="comment.image"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
const comments = ref<Comment[]>([]);

// define type comment
interface Comment {
  id: number;
  text: string;
  author: string;
  date: string;
  profileImage: string;
  image: string;
}

const likes = ref(0);
const dislikes = ref(0);
const showComments = ref(false);
const commentText = ref("");
const hover = ref(false);
const showModal = ref(false);
const user = ref("Anonymous");
const createdAt = ref(new Date().toLocaleString());
const title = ref("Post Title");
const content = ref("Post Content");
const image = ref("");

const incrementLikes = () => likes.value++;

const incrementDislikes = () => dislikes.value++;

const createComment = () => {
  comments.value.push({
    id: comments.value.length + 1,
    text: commentText.value,
    author: "Anonymous",
    date: new Date().toLocaleString(),
    profileImage: "https://i.pravatar.cc/100",
    image: "",
  });
  commentText.value = "";
};
</script>

<!-- <script>
export default {
  props: {
    title: String,
    content: String,
    user: String,
    createdAt: String,
    image: String,
    comments: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      likes: 0,
      dislikes: 0,
      showComments: false,
    };
  },
  methods: {
    // For testing purposes
    incrementLikes() {
      this.likes++;
    },
    incrementDislikes() {
      this.dislikes++;
    },
  },
};
</script> -->

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");

.post {
  margin-bottom: 20px;
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 4px;
  background-color: #f5f5f5;
  font-family: "Poppins", sans-serif;
}

.post-title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 10px;
  word-wrap: break-word;
}

.post-content {
  font-size: 16px;
  line-height: 1.5;
  margin-bottom: 20px;
  word-wrap: break-word;
}

.post-meta {
  color: #777;
  font-size: 14px;
  margin-bottom: 10px;
}

.profile-and-button {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.follow-button {
  background-color: #007bff;
  border: none;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  padding: 5px 10px;
  margin-left: 15px;
}

.follow-button:hover {
  background-color: #0056b3;
}

.post-buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
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

.post:hover {
  transform: scale(1.02);
  transition: transform 0.3s ease-in-out;
}

.show-comments-button {
  background-color: #007bff;
  border: none;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  padding: 5px 10px;
  margin-top: 10px;
}

.post-comments {
  margin-top: 10px;
}

.no-comments {
  margin-top: 10px;
  font-size: 14px;
  color: #888;
}

.create-comment {
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
}

.comment-textarea {
  resize: none;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 10px;
  font-family: "Poppins", sans-serif;
  font-size: 14px;
  line-height: 1.5;
  margin-bottom: 10px;
}

.comment-button {
  background-color: #007bff;
  border: none;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  padding: 8px 15px;
  border-radius: 4px;
}

.comment-button:hover {
  background-color: #0056b3;
}
</style>
