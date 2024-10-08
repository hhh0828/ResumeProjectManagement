document.addEventListener("DOMContentLoaded", function () {
    let allprojectsData = [];
    let currentIndex = 0;
    const projectsPerPage = 2; // 한 번에 표시할 프로젝트의 수

    // 컨테이너와 버튼 요소 가져오기
    const projectContainer = document.getElementById('projectContainer');
    const loadMoreButton = document.getElementById('loadMoreButton');

    // 프로젝트 데이터를 서버에서 가져오기
    fetch('/returnproject')
        .then(response => response.json())
        .then(data => {
            allprojectsData = data.projects;
            displayProjects(); // 페이지 로드 시 초기 두 개의 프로젝트만 표시
        })
        .catch(error => console.error('Error fetching data:', error));

    // 프로젝트를 표시하는 함수
    function displayProjects() {
        const endIndex = Math.min(currentIndex + projectsPerPage, allprojectsData.length);

        for (let i = currentIndex; i < endIndex; i++) {
            const project = allprojectsData[i];

            const projectCard = `
                <div class="card overflow-hidden shadow rounded-4 border-0 mb-5">
                        
                    <div class="card-body p-0">
                        <div class="d-flex align-items-center project-card">
                            <div class="text-content p-5">
                                <input type="hidden" name="post_id" value="${project.ID}}" data-id=${project.ID}}>
                                <h2 class="fw-bolder">${project.Name}</h2>
                                <p>${project.shortdesc}</p>
                                <a href="${project.detailurl}" class="stretched-link"></a>
                            </div>
                            <div class="image-container">
                            <img class="img-fluid project-image" src="${project.imgurl}" alt="${project.Name}" />
                            </div>
                        </div>
                        
                    </div>
                    <div>
                <button class="btn btn-light position-absolute top-0 start-0 m-2 p-1 edit-button" 
                style="border-radius: 50%; z-index: 10;"
                data-id="${project.ID}">
                <i class="bi bi-pencil"></i>
                </button>
                </div>
                </div>
                
            `;

            projectContainer.insertAdjacentHTML('beforeend', projectCard);
        }

        currentIndex = endIndex;

        // 모든 프로젝트를 표시했으면 "더보기" 버튼 숨기기
        if (currentIndex >= allprojectsData.length) {
            loadMoreButton.style.display = 'none';
        }
    }
    // "더보기" 버튼 클릭 이벤트
    loadMoreButton.addEventListener('click', function () {
        displayProjects(); // 버튼 클릭 시 추가 프로젝트 로드
    });
});
document.addEventListener('DOMContentLoaded', function() {
    //  .edit-button인지 확인
    document.body.addEventListener('click', function(event) {
        const button = event.target.closest('.edit-button');
        if (button) {
            const projectId = button.getAttribute('data-id'); // 프로젝트 ID 가져오기
            console.log('Project ID:', projectId); // 프로젝트 ID 출력
                fetch(`/editproject?id=${projectId}`) // 템플릿 리터럴로 변수 삽입
                 .then(response => {
                  if (response.ok) {
                        alert(`${projectId}`+'로 이동합니다')
                        window.location.href = `/editproject?id=${projectId}`;
                 } else if (response.status === 401) {
                       // 401 에러가 발생하면 로그인 페이지로 리디렉션
                       alert('권한이 없습니다, 해당 작업은 로그인이 필요합니다, 로그인 페이지로 이동합니다')
                         window.location.href = '/loginpage'; // 로그인 페이지로 이동
                 } else {
                     // 다른 에러 처리
                      console.error('Error:', response.status);
                   }
               })
               .catch(error => {
           // 네트워크 오류 처리
                  console.error('Network error:', error);
                   });

              }
        
    });
});