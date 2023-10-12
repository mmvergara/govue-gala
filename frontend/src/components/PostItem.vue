<template>
  <section
    class="bg-white p-2 border-t-[purple] border-t-4 w-full max-w-[500px]"
  >
    <h2>
      {{ post.post_title }}
    </h2>
    <p>{{ post.post_description }}</p>
    <span class="opacity-50 ml-auto text-xs">{{ time }}</span>
  </section>
</template>

<script setup lang="ts">
import { defineProps } from "vue";
import type { PostType } from "../types/post-types";
const { post } = defineProps<{ post: PostType }>();

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
