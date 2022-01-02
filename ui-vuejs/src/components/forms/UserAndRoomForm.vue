<template>

  <q-form
    class="column serverForm"
  >
    <q-chip class="label" square>
      <q-avatar color="red" icon="bookmark" text-color="white"/>
      Add Users or Rooms
    </q-chip>
    <div class="row content-stretch formRow">
      <AddUser class="col"/>

    </div>
    <div class="row content-stretch formRow">
      <AddRoom class="col"/>
    </div>
  </q-form>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import AddUser from 'components/forms/AddUser.vue';
import AddRoom from 'components/forms/AddRoom.vue';


export default defineComponent({
  name: 'UserAndRoomForm',
  components: {AddRoom, AddUser},
  setup() {
    return {}
  }
})
</script>

<style lang="sass" scoped>

.formRow
  padding-left: 10px
  padding-right: 10px


.spacer
  width: 10px

.submitBtn
  margin-top: 10px

</style>
