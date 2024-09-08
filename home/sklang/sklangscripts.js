document.addEventListener("DOMContentLoaded", function () {
    let allExperienceData = [];
    let currentIndex = 0;
    const experienceContainer = document.getElementById('experienceContainer');
    const loadMoreButton = document.getElementById('loadMoreButton');
    const experiencesPerPage = 2; // 한 번에 표시할 경험의 수

    // API로부터 JSON 데이터 가져오기
    fetch('/returnresume')
        .then(response => response.json())
        .then(data => {
            allExperienceData = data.exps;
            displayExperience(); // 페이지 로드 시 초기 두 개의 경험만 표시

            function displayExperience() {
                const endIndex = Math.min(currentIndex + experiencesPerPage, allExperienceData.length);

                for (let i = currentIndex; i < endIndex; i++) {
                    const experience = allExperienceData[i];

                    const experienceCard = ` 
                        <div class="card shadow border-0 rounded-4 mb-5" data-id="${experience.ID}">
                            <div class="card-body p-5">
                                <div class="row align-items-center gx-5">
                                    <div class="col text-center text-lg-start mb-4 mb-lg-0">
                                        <div class="bg-light p-4 rounded-4">
                                            <div class="text-primary fw-bolder mb-2">${experience.period}</div>
                                            <div class="small fw-bolder">${experience.role}</div>
                                            <div class="small text-muted">${experience.company}</div>
                                        </div>
                                    </div>
                                    <div class="col-lg-8">
                                        <div>${experience.description}</div>
                                    </div>
                                </div>
                            </div>
                            <button class="btn btn-light position-absolute top-0 start-0 m-2 p-1 edit-button" 
                            style="border-radius: 50%; z-index: 10;">
                            <i class="bi bi-pencil"></i>
                            </button>
                        </div>
                    `;

                    experienceContainer.insertAdjacentHTML('beforeend', experienceCard);
                }

                currentIndex = endIndex;

                // 모든 경험을 다 표시했으면 "더보기" 버튼 숨기기
                if (currentIndex >= allExperienceData.length) {
                    loadMoreButton.style.display = 'none';
                }
            }

            loadMoreButton.addEventListener('click', function () {
                displayExperience(); // 버튼 클릭 시 추가 데이터 로드
            });
        })
        .catch(error => console.error('Error fetching data:', error));

    // 업로드 버튼 클릭 시 새로운 카드 열기
    document.getElementById('upload').addEventListener('click', function () {
        const uniqueUploadId = `uploadCard-${Date.now()}`;
        const newCard = `
            <div class="card shadow border-0 rounded-4 mb-5" id="${uniqueUploadId}">
                <div class="card-body p-5">
                    <h3 class="fw-bolder mb-4">Add New Experience</h3>
                    <form id="uploadForm-${uniqueUploadId}">
                        <div class="mb-3">
                            <label for="periodStart" class="form-label">Start Date</label>
                            <input type="month" class="form-control" id="periodStart-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="periodEnd" class="form-label">End Date</label>
                            <input type="month" class="form-control" id="periodEnd-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="role" class="form-label">Role</label>
                            <input type="text" class="form-control" id="role-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="company" class="form-label">Company</label>
                            <input type="text" class="form-control" id="company-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="description" class="form-label">Description</label>
                            <textarea class="form-control" id="description-${uniqueUploadId}" rows="3" required></textarea>
                        </div>
                        <button type="submit" class="btn btn-primary">Submit</button>
                        <button type="button" id="cancelUpload-${uniqueUploadId}" class="btn btn-secondary">Cancel</button>
                    </form>
                </div>
            </div>
        `;
    
        experienceContainer.insertAdjacentHTML('afterbegin', newCard); // 카드를 가장 위로 추가
    
        document.getElementById(`uploadForm-${uniqueUploadId}`).addEventListener('submit', function (event) {
            event.preventDefault(); // 폼 제출 기본 동작 방지
    
            const periodStart = document.getElementById(`periodStart-${uniqueUploadId}`).value;
            const periodEnd = document.getElementById(`periodEnd-${uniqueUploadId}`).value;
            const role = document.getElementById(`role-${uniqueUploadId}`).value;
            const company = document.getElementById(`company-${uniqueUploadId}`).value;
            const description = document.getElementById(`description-${uniqueUploadId}`).value;
    
            // YYYY-MM 형식으로 입력된 날짜를 년-월 형식으로 변환
            const formatPeriod = (start, end) => {
                const [startYear, startMonth] = start.split("-");
                const [endYear, endMonth] = end.split("-");
                return `${startYear}년 ${startMonth}월 ~ ${endYear}년 ${endMonth}월`;
            };
    
            const periodFormatted = formatPeriod(periodStart, periodEnd);
    
            fetch('/uploadresume', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    period: periodFormatted,
                    role,
                    company,
                    description
                })
            })
            .then(response => response.json())
            .then(result => {
                alert('Experience added successfully!');
                document.getElementById(uniqueUploadId).remove(); // 새 카드 제거
            })
            .catch(error => console.error('Error:', error));
        });
    
        document.getElementById(`cancelUpload-${uniqueUploadId}`).addEventListener('click', function () {
            document.getElementById(uniqueUploadId).remove(); // 새 카드 제거
        });
    });

    // 수정 버튼 클릭 시 카드 내에 수정 폼 추가하기
    experienceContainer.addEventListener('click', function (event) {
        if (event.target.closest('.edit-button')) {
            const card = event.target.closest('.card');
            const cardId = card.getAttribute('data-id');
    
            // 카드에서 기존 기간 데이터를 추출 (예: "2024년 12월 ~ 2025년 1월")
            const periodText = card.querySelector('.text-primary').textContent;
            const [startPeriod, endPeriod] = periodText.split(' ~ ').map(period => {
                const [year, month] = period.split('년 ')[0].split(' ');
                return `${year}-${month.replace('월', '').padStart(2, '0')}`; // YYYY-MM 형식으로 변환
            });
    
            // 수정 카드 추가
            const editCardId = `editCard-${Date.now()}`;
            const editCard = `
                <div class="card shadow border-0 rounded-4 mb-5" id="${editCardId}">
                    <div class="card-body p-5">
                        <h3 class="fw-bolder mb-4">Edit Experience</h3>
                        <form id="editForm-${editCardId}">
                            <input type="hidden" id="editId" value="${cardId}">
                            <div class="mb-3">
                                <label for="editPeriodStart-${editCardId}" class="form-label">Start Date</label>
                                <input type="month" class="form-control" id="editPeriodStart-${editCardId}" value="${startPeriod}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editPeriodEnd-${editCardId}" class="form-label">End Date</label>
                                <input type="month" class="form-control" id="editPeriodEnd-${editCardId}" value="${endPeriod}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editRole-${editCardId}" class="form-label">Role</label>
                                <input type="text" class="form-control" id="editRole-${editCardId}" value="${card.querySelector('.small.fw-bolder').textContent}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editCompany-${editCardId}" class="form-label">Company</label>
                                <input type="text" class="form-control" id="editCompany-${editCardId}" value="${card.querySelector('.small.text-muted').textContent}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editDescription-${editCardId}" class="form-label">Description</label>
                                <textarea class="form-control" id="editDescription-${editCardId}" rows="3" required>${card.querySelector('.col-lg-8 div').textContent}</textarea>
                            </div>
                            <button type="submit" class="btn btn-primary">Save Changes</button>
                            <button type="button" id="cancelEdit-${editCardId}" class="btn btn-secondary">Cancel</button>
                        </form>
                    </div>
                </div>
            `;
    
            card.insertAdjacentHTML('beforeend', editCard);
    
            document.getElementById(`editForm-${editCardId}`).addEventListener('submit', function (event) {
                event.preventDefault(); // 폼 제출 기본 동작 방지
    
                const periodStart = document.getElementById(`editPeriodStart-${editCardId}`).value;
                const periodEnd = document.getElementById(`editPeriodEnd-${editCardId}`).value;
                const role = document.getElementById(`editRole-${editCardId}`).value;
                const company = document.getElementById(`editCompany-${editCardId}`).value;
                const description = document.getElementById(`editDescription-${editCardId}`).value;
    
                const formatPeriod = (start, end) => {
                    const [startYear, startMonth] = start.split("-");
                    const [endYear, endMonth] = end.split("-");
                    return `${startYear}년 ${startMonth}월 ~ ${endYear}년 ${endMonth}월`;
                };
    
                const periodFormatted = formatPeriod(periodStart, periodEnd);
    
                fetch('/editresume', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        id: cardId,
                        period: periodFormatted,
                        role,
                        company,
                        description
                    })
                })
                .then(response => response.json())
                .then(result => {
                    alert('Experience updated successfully!');
                    document.getElementById(editCardId).remove(); // 수정 카드 제거
                    // 여기에서 기존 카드 내용 업데이트 필요 (예: 카드의 기간, 역할, 회사 정보)
                    card.querySelector('.text-primary').textContent = periodFormatted;
                    card.querySelector('.small.fw-bolder').textContent = role;
                    card.querySelector('.small.text-muted').textContent = company;
                    card.querySelector('.col-lg-8 div').textContent = description;
                })
                .catch(error => console.error('Error:', error));
            });
    
            document.getElementById(`cancelEdit-${editCardId}`).addEventListener('click', function () {
                document.getElementById(editCardId).remove(); // 수정 카드 제거
            });
        }
    });
});
