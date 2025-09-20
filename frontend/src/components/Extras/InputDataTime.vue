<script setup lang="ts">

interface Props {
  id: string,
  name: string
}

defineProps<{
  props: Props
}>()

</script>

<template>
  <div style="padding: 0.5rem;">
    <div style="padding: 0.5rem;">
      <label style="font-size: var(--font-size);" :for="props.id">{{ props.name }}</label>
    </div>
    <div style="display: flex; align-items: center;">
      <slot></slot>
    </div>
  </div>
</template>
