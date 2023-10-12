<script setup lang="ts">
import PostItem from "@/components/PostItem.vue";
import { ref } from "vue";
import { AxiosRequest } from "@/components/AxiosInstance";
import type { PostType } from "@/types/post-types";

const posts = ref<PostType[] | null>(null);

const getPosts = async () => {
  const { data } = await AxiosRequest<{ posts: PostType[] }>("/post");
  posts.value = data.posts;
};

const removePost = async (id: string) => {
  await AxiosRequest(`/post/${id}`, "DELETE");
  posts.value = posts.value?.filter((post) => post.post_id !== id) || null;
};
getPosts();
</script>

<template>
  <main>
    <h1 class="text-center py-8 text-[50px]">Govue Gala</h1>
    <section class="flex justify-center items-center flex-col gap-4">
      <PostItem
        v-for="post in posts"
        :post="post"
        :key="post.created_at"
        :remove-post="removePost"
      />
    </section>
  </main>
</template>
