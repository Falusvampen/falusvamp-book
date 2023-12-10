<template>
  <div class="comment">
    <div class="profile-and-button">
      <profileImage :image="profileImage" />
      <button class="follow-button">Follow</button>
    </div>
    <div class="comment-meta">
      <p class="comment-author">Comment by: {{ author }}</p>
      <p class="comment-date">Created at: {{ date }}</p>
    </div>
    <div class="comment-content">
      <img :src="image" class="post-image" style="width: 100%" />
      <p class="comment-text">{{ text }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    text: String,
    author: String,
    date: String,
    profileImage: String,
    image: String,
  },
};
</script>

<style scoped>
.comment {
  /* display: inline-block; */
  width: 100%;
  /* margin-bottom: 20px; */
  /* background-color: #d5d5d5; */
  border-radius: 4px;
  border: 1px solid #ccc;
  margin-top: 5px;
  /* padding-top: 5px; */
  /* padding: 10px; */
  /* margin: 10px; */
}

.profile-and-button {
  display: flex;
  align-items: center;
  /* margin-bottom: 10px; */
  margin: 10px;
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

.profile-image img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.comment-content {
  flex: 1;
  margin: 10px;
}

.comment-text {
  font-size: 16px;
  margin-bottom: 5px;
  word-break: break-word;
}

.comment-meta {
  /* display: flex; */
  /* justify-content: space-between; */
  color: #777;
  font-size: 14px;
  margin: 10px;
}

.comment-author {
  margin-right: 10px;
}

.comment-date {
  color: #888;
}
</style>
