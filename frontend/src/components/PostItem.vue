<template>
  <section
    class="bg-white p-4 border-[purple] border-t-2 drop-shadow-md rounded-md w-full max-w-[500px]"
  >
    <h2 class="font-semibold text-lg">
      {{ post.post_title }}
    </h2>
    <p>{{ post.post_description }}</p>
    <div class="w-[99%] border-b-[1px] mx-auto my-4"></div>
    <div class="flex justify-between items-center">
      <span class="opacity-50 text-xs">{{ time }}</span>
      <button
        class="bg-white border-2 hover:bg-slate-300 text-sm p-2 rounded-sm"
        @click="handleDelete"
      >
        ❌
      </button>
    </div>
  </section>
</template>

<script setup lang="ts">
import { defineProps } from "vue";
import type { PostType } from "../types/post-types";
const { post, removePost } = defineProps<{
  post: PostType;
  removePost: (id: string) => void;
}>();

const handleDelete = () => {
  removePost(post.post_id);
};

const timeFromNow = (date: string) => {
  const time = new Date(date);
  const now = new Date();
  const diff = now.getTime() - time.getTime();
  const seconds = Math.floor(diff / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  if (days > 0) {
    return `${days} days ago`;
  } else if (hours > 0) {
    return `${hours} hours ago`;
  } else if (minutes > 0) {
    return `${minutes} minutes ago`;
  } else {
    return `${seconds} seconds ago`;
  }
};
const time = timeFromNow(post.created_at);
</script>
