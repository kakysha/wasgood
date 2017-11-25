jQuery(document).ready(function ($) {
	$('.popup').popup({transition: 'all 0.3s'});

	$('.rating .btn').click(function(){
		if (typeof $('body').data('user-id') === "undefined") {
			$('.popup').popup('show');
			return;
		}
		var $btn = $(this);
		var $parent = $btn.parents('.item');
		var item_id = $parent.data('item-id');
		var voice = $btn.parent().hasClass('plus') ? 1 : -1;
		// next we will handle not only this item but all items on page having same id
		$parent_rating = $('.item[data-item-id='+item_id+'] .rating');
		$btn = $parent_rating.find('.' + (voice>0?'plus':'minus') + ' .btn');
		$btn.attr('disabled', 'disabled');
		var $opposite_btn = $parent_rating.find('.btn').not($btn);
		var $opposite_count = $opposite_btn.siblings('.count'), opposite_count = parseInt($opposite_count.first().text());
		var voted = typeof $opposite_btn.attr('disabled') !== "undefined";
		var $total = $parent_rating.find('.total'), total = parseInt($total.first().text());
		var $count = $btn.siblings('.count'), count = parseInt($count.first().text());
		$.ajax({
			url: $parent_rating.data('vote-url')+voice,
			type: 'POST',
			success: function(result) {
				if (voted) {
					$opposite_btn.removeAttr('disabled');
					$total.text(total+2*voice);
					$opposite_count.text(opposite_count-1);
				} else {
					$total.text(total+voice);
				}
				$count.text(count+1);
			},
			error: function(xhr, status, response) {
				alert(response);
			}
		});
		publish_vote($parent.find('.brand').text(), $parent.find('.name').text(), voice, $parent.find('.name').attr('href'));
	});

	$('.review .delete').click(function(e){
		e.preventDefault();
		$btn = $(this);
		$review = $btn.parents('.review');
		$btn.detach();
		$.ajax({
			url: $btn.attr('href'),
			type: 'DELETE',
			success: function(result) {
				$review.remove();
			},
			error: function(xhr, status, response) {
				alert(response);
			}
		});
	});

	$('.review .approve').click(function(e){
		e.preventDefault();
		$btn = $(this);
		$review = $btn.parents('.review');
		$btn.detach();
		$.ajax({
			url: $btn.attr('href'),
			type: 'POST',
			success: function(result) {
				$review.removeClass('unapproved');
			},
			error: function(xhr, status, response) {
				alert(response);
			}
		});
	});

	$('.pagination button').click(function(e){
		var pagenum = $(this).parents('.pagination').data('pagenum');
		var targetpage = $(this).hasClass('prev') ? pagenum-1 : pagenum+1;
		var url = location.href;
		location.href = updateUrlParameter(url, "page", targetpage);
	});

	$("[data-active]").each(function(){
		var $el = $(this);
		var targetURL = $el.find('a').addBack('a').first().attr('href');
		var url = location.href;
		if (url.indexOf(targetURL) > -1) {
			$el.addClass('uk-active');
		}
	});
	if (typeof posttowall === "undefined") {
		posttowall = function(text) {} // will be overriden in footer by social network api js
	}
	function publish_vote(brand, name, side, url){
		if (!$("input[name='sharing']").prop("checked"))
			return;
		posttowall(brand.trim() + " " + name.trim() + "\n" + " получает от меня " + (side > 0 ? '+1' : '-1') + " в рейтинге жидкостей " + document.location.host, url);
	}

	$("input[type='checkbox'].local").each(function() {
		var mycookie = Cookies.get($(this).attr('name'));
		if (typeof mycookie !== "undefined") {
			$(this).prop('checked', mycookie == "true");
		}
	});
	$("input[type='checkbox'].local").change(function() {
		Cookies.set($(this).attr("name"), $(this).prop('checked'), { expires: 60 });
	});

	var searchInProgress = null;
	var old_q = "";
	var $result_list_items = [];
	var selected_item_idx = -1;
	$('.uk-search-field[name=q]').on('keyup click', function(e) {
		if ($('.ajax-search-results').is(':visible')) {
			if (e.which == 40) { // down
				$result_list_items.eq(selected_item_idx++).removeClass('selected');
				if (selected_item_idx == $result_list_items.length)
					selected_item_idx = 0;
			}
			if (e.which == 38) { // up
				$result_list_items.eq(selected_item_idx--).removeClass('selected');
				if (selected_item_idx < 0)
					selected_item_idx = $result_list_items.length-1;
			}
			if (e.which == 38 || e.which == 40) {
				$result_list_items.eq(selected_item_idx).addClass('selected');
			}
		}

		var $dropdown = UIkit.dropdown($('.ajax-search-results').parents('[data-uk-dropdown]'));

		var q = $(this).val();
		if (q.length < 3) {
			$dropdown.hide();
			return;
		}
		if (q == old_q) {
			$dropdown.show();
			return;
		}

		if (searchInProgress && searchInProgress.readyState != 4) searchInProgress.abort();

		searchInProgress = $.get("/search", {q: q, ajax: 1}, function(html){
			$('.ajax-search-results').html(html);
			$dropdown.show();
			$result_list_items = $('.ajax-search-results li');
			selected_item_idx = -1;
			old_q = q;
		});
	}).parents('form').submit(function(e) {
		if (selected_item_idx != -1) {
			e.preventDefault();
			$result_list_items.eq(selected_item_idx).find('a')[0].click();
		}
	});
});

// Add / Update a key-value pair in the URL query parameters
function updateUrlParameter(uri, key, value) {
	// remove the hash part before operating on the uri
	var i = uri.indexOf('#');
	var hash = i === -1 ? ''  : uri.substr(i);
		 uri = i === -1 ? uri : uri.substr(0, i);

	var re = new RegExp("([?&])" + key + "=.*?(&|$)", "i");
	var separator = uri.indexOf('?') !== -1 ? "&" : "?";
	if (uri.match(re)) {
		uri = uri.replace(re, '$1' + key + "=" + value + '$2');
	} else {
		uri = uri + separator + key + "=" + value;
	}
	return uri + hash;  // finally append the hash as well
}