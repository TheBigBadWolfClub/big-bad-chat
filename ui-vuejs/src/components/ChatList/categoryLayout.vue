<template>
  <q-expansion-item
    :caption=model.caption
    :icon=model.icon
    :label=model.label
    default-opened
  >
    <q-card>
      <q-card-section>
        <slot></slot>
      </q-card-section>
    </q-card>
  </q-expansion-item>
</template>

<script lang="ts">


export interface CategoryOptions {
  caption: string,
  icon: string,
  label: string
}

import {defineComponent, PropType} from 'vue';

export default defineComponent({
  name: 'CategoryLayout',
  components: {},
  props: {
    model: {
      type: Object as PropType<CategoryOptions>,
      required: true
    }

  }
})
</script>

<style lang="sass">
.entry
  width: 100%

</style>

