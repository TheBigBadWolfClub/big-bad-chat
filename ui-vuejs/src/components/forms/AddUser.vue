<template>

  <q-input
    v-model="username"
    :rules="[ val => val && val.length > 0 ]"
    bg-color="grey-3"
    bottom-slots
    color="grey-10"
    dense
    label="Add User"
    label-color="black"
    lazy-rules
    outlined>


    <template v-slot:after>
      <q-btn
        color="secondary" dense
        icon="add"
        square
        type="submit"
        @click="onSubmit"
      />
    </template>
  </q-input>

</template>

<script lang="ts">
import {defineComponent, ref} from 'vue';
import {useQuasar} from 'quasar';
import axios from 'axios';


export default defineComponent({
  name: 'AddUser',
  components: {},
  setup() {
    const $q = useQuasar();

    const onSubmit = (() => {
      const user = {
        name: usernameRef.value
      }

      axios.post('/users', user)
        .then(() => {
          $q.notify('User  Created')
        })
        .catch((err) => $q.notify(`Fail to create User : ${err}`))
    })

    let usernameRef = ref('')
    return {onSubmit, username: usernameRef}
  }
})
</script>
