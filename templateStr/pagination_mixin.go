package templateStr

var Pagination  = `
import {pageList} from '@/api/user.js'
import {resetArgs} from "@/utils/index.js";

export default {
    data() {
        return {
            requestFunc:pageList,
            searchParams: {
                postTitle: '',
                postType: '',
                postStatus: '',
            },
			tableData: []
        }
    },
    methods: {
        returnData(pageList){
            this.tableData = pageList.list
        },
        clearSearchParams(){
            resetArgs(this.searchParams);
            this.refresh();
        },
        refresh(){
            this.$refs.pagination.Refresh()
        }
    },
}
`
